package service

import (
	"context"
	"go-fiber-api/internal/config"
	"go-fiber-api/internal/model"
	"go-fiber-api/internal/repository"
	"go-fiber-api/pkg/dto"
	"go-fiber-api/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo    repository.UserRepository
	redisClient *redis.Client
	config      *config.Config
}

func NewUserService(userRepo repository.UserRepository, redisClient *redis.Client, config *config.Config) *UserService {
	return &UserService{
		userRepo:    userRepo,
		redisClient: redisClient,
		config:      config,
	}
}
func (s *UserService) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := s.userRepo.FindOne(ctx, bson.M{"email": email})
	if err != nil || user == nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Count(ctx context.Context, query bson.D) (int64, error) {
	return s.userRepo.Count(ctx, query)
}

func (s *UserService) FindAll(ctx context.Context, query bson.D, opts *options.FindOptions) ([]model.User, error) {
	return s.userRepo.FindAll(ctx, query, opts)
}

func (s *UserService) Create(ctx context.Context, payload *dto.RegisterRequest) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if len(payload.Roles) == 0 {
		payload.Roles = []string{string(utils.UserRole)}
	}

	now := time.Now()
	user := &model.User{
		ID:        primitive.NewObjectID(),
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  string(hashedPassword),
		Roles:     payload.Roles,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Login(ctx context.Context, password string, user *model.User) (*utils.TokenPair, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	auth := utils.NewAuthHandler(s.config.JWTSecretKey, s.config.JWTRefreshKey, s.config.JWTExpiresIn, s.config.JWTRefreshIn)
	tokenPair, err := auth.GenerateTokenPair(user.ID.Hex(), user.Roles)
	if err != nil {
		return nil, err
	}

	expires, _ := time.ParseDuration(s.config.JWTExpiresIn)
	if err := s.redisClient.Set(ctx,
		tokenPair.AccessToken,
		user.ID.Hex(),
		expires).Err(); err != nil {
		return nil, err
	}
	return tokenPair, nil
}

func (s *UserService) RefreshToken(ctx context.Context, refreshToken string) (*utils.TokenPair, error) {
	auth := utils.NewAuthHandler(s.config.JWTSecretKey, s.config.JWTRefreshKey, s.config.JWTExpiresIn, s.config.JWTRefreshIn)
	claims, err := auth.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	blacklisted, err := s.redisClient.Get(ctx, "blacklist:"+refreshToken).Result()
	if blacklisted != "" || err != redis.Nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Refresh token has been revoked")
	}

	user, err := s.FindByID(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}

	tokenPair, err := auth.GenerateTokenPair(user.ID.Hex(), user.Roles)
	if err != nil {
		return nil, err
	}

	expires, _ := time.ParseDuration(s.config.JWTExpiresIn)
	if err := s.redisClient.Set(ctx,
		tokenPair.AccessToken,
		user.ID.Hex(),
		expires).Err(); err != nil {
		return nil, err
	}

	// Keep blacklist for 48h to prevent reuse
	if err := s.redisClient.Set(ctx,
		"blacklist:"+refreshToken,
		"true",
		48*time.Hour).Err(); err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func (s *UserService) FindByID(ctx context.Context, id string) (*model.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return s.userRepo.FindOne(ctx, bson.M{"_id": objID})
}

func (s *UserService) UpdateById(ctx context.Context, id primitive.ObjectID, payload *dto.UpdateUserRequest) error {
	_, err := s.userRepo.UpdateByID(ctx, id, payload)
	return err
}

func (s *UserService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.userRepo.Delete(ctx, id)
}

// check redis for blacklisted token
func (s *UserService) ValidateTokenWithRedis(ctx context.Context, token string) error {
	// Check blacklist with 24h window
	blacklisted, err := s.redisClient.Get(ctx, "blacklist:"+token).Result()
	if err != redis.Nil || blacklisted != "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Token has been revoked")
	}

	// Check active tokens
	_, err = s.redisClient.Get(ctx, token).Result()
	if err == redis.Nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Token not found in active sessions")
	}
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Logout(ctx context.Context, accessToken, refreshToken string) error {
	pipe := s.redisClient.Pipeline()

	// Blacklist access token for 24h
	expires, _ := time.ParseDuration(s.config.JWTExpiresIn)
	pipe.Set(ctx,
		"blacklist:"+accessToken,
		"true",
		expires)

	// Blacklist refresh token for 48h
	pipe.Set(ctx,
		"blacklist:"+refreshToken,
		"true",
		48*time.Hour)

	// Remove active access token
	pipe.Del(ctx, accessToken)

	_, err := pipe.Exec(ctx)
	return err
}

package handlers

import (
	"context"
	"go-fiber-api/internal/model"
	"go-fiber-api/internal/service"
	"go-fiber-api/pkg/dto"
	"go-fiber-api/pkg/middleware"
	"go-fiber-api/pkg/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// @Summary List users
// @Description Get paginated list of users with optional filtering
// @Tags admin
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Param name query string false "Filter by name"
// @Success 200
// @Router /admin/users [get]
func (u *UserHandler) UserList(c *fiber.Ctx) error {
	page, pageSize := utils.PaginationParams(c)

	var filter dto.UserFilter
	if err := c.QueryParser(&filter); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid filter parameters")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mongoFilter := bson.D{}
	if filter.Name != "" {
		mongoFilter = append(mongoFilter, bson.E{
			Key: "name",
			Value: bson.D{{
				Key:   "$regex",
				Value: primitive.Regex{Pattern: filter.Name, Options: "i"},
			}},
		})
	}

	total, err := u.userService.Count(ctx, mongoFilter)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to count users: "+err.Error())
	}

	opts := options.Find().
		SetSkip(int64((page - 1) * pageSize)).
		SetLimit(int64(pageSize)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	users, err := u.userService.FindAll(ctx, mongoFilter, opts)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, err.Error())
	}

	response := utils.CreatePagination(page, pageSize, total, users)
	return utils.SendSuccess(c, http.StatusOK, response)
}

// @Summary Register endpoint
// @Description Post the API's register
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "User registration details"
// @Router /auth/register [post]
func (u *UserHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return utils.SendValidationError(c, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	existingUser, err := u.userService.FindByEmail(ctx, req.Email)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, err.Error())
	}

	if existingUser != nil {
		return utils.SendError(c, http.StatusBadRequest, "Email already exists")
	}

	user, err := u.userService.Create(ctx, &req)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, err.Error())
	}

	token, err := u.userService.Login(ctx, req.Password, user)
	if err != nil {
		return utils.SendError(c, http.StatusUnauthorized, "Invalid password")
	}

	info := &model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Roles: user.Roles,
	}

	res := fiber.Map{
		"info":  info,
		"token": token,
	}
	return utils.SendSuccess(c, http.StatusOK, res)
}

// @Summary Login endpoint
// @Description Post the API's login
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "User login"
// @Router /auth/login [post]
func (u *UserHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return utils.SendValidationError(c, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := u.userService.FindByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return utils.SendError(c, http.StatusNotFound, "Invalid email")
	}

	tokenPair, err := u.userService.Login(ctx, req.Password, user)
	if err != nil {
		return utils.SendError(c, http.StatusUnauthorized, "Invalid password")
	}
	return utils.SendSuccess(c, http.StatusOK, tokenPair, "Login successful")
}

// @Summary Refresh endpoint
// @Description Post the API's refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body dto.RefreshTokenRequest true "Refresh token"
// @Router /auth/refresh [post]
func (u *UserHandler) RefreshToken(c *fiber.Ctx) error {
	var req dto.RefreshTokenRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return utils.SendValidationError(c, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tokenPair, err := u.userService.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return utils.SendError(c, http.StatusUnauthorized, err.Error())
	}

	if tokenPair == nil {
		return utils.SendError(c, http.StatusUnauthorized, "Invalid refresh token")
	}

	return utils.SendSuccess(c, http.StatusOK, tokenPair, "Refresh token successful")
}

// @Summary Profile endpoint
// @Description Get the API's get profile
// @Tags user
// @Accept json
// @Produce json
// @Security Bearer
// @Router /user/profile [get]
func (u *UserHandler) GetProfile(c *fiber.Ctx) error {
	user, ok := middleware.GetUserFromContext(c)
	if !ok {
		return utils.SendError(c, http.StatusUnauthorized, "User not found")
	}

	res := &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Roles:     user.Roles,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return utils.SendSuccess(c, http.StatusOK, res)
}

// @Summary Update endpoint
// @Description Get the API's update user
// @Tags user
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "User ID"
// @Param request body dto.UpdateUserRequest true "User update details"
// @Router /admin/user/{id} [put]
func (u *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var req dto.UpdateUserRequest
	id := c.Params("id")

	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := utils.ValidateStruct(&req); err != nil {
		return utils.SendValidationError(c, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid ID format")
	}

	user, err := u.userService.FindByID(ctx, objID.Hex())
	if err != nil || user == nil {
		return utils.SendError(c, http.StatusNotFound, "User not found")
	}

	if err := u.userService.UpdateById(ctx, objID, &req); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, err.Error())
	}

	res := fiber.Map{
		"id":   user.ID,
		"name": req.Name,
	}

	return utils.SendSuccess(c, http.StatusOK, res, "Profile updated successfully")
}

// @Summary Delete endpoint
// @Description Get the API's delete user
// @Tags admin
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "User ID"
// @Router /admin/user/{id} [delete]
func (u *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	paramId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid user ID format")
	}

	auth, ok := middleware.GetUserFromContext(c)
	if !ok {
		return utils.SendError(c, http.StatusUnauthorized, "Invalid session")
	}

	user, err := u.userService.FindByID(ctx, paramId.Hex())
	if err != nil || user == nil {
		return utils.SendError(c, http.StatusNotFound, "User not found")
	}

	if user.ID == auth.ID {
		return utils.SendError(c, http.StatusUnauthorized, "You cannot delete yourself")
	}

	if err := u.userService.Delete(ctx, paramId); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, nil, "User deleted successfully")
}

// @Summary Logout endpoint
// @Description Post the API's logout
// @Tags auth
// @Accept json
// @Produce json
// @Security Bearer
// @Header 200 {string} X-Refresh-Token "Refresh token for logout"
// @Param X-Refresh-Token header string true "Refresh token"
// @Router /auth/logout [get]
func (u *UserHandler) Logout(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get access token from context
	token, ok := c.Locals("token").(string)
	if !ok {
		return utils.SendError(c, http.StatusUnauthorized, "Token not found")
	}

	// Get refresh token from header
	refreshToken := c.Get("X-Refresh-Token")
	if refreshToken == "" {
		return utils.SendError(c, http.StatusBadRequest, "Refresh token is required")
	}

	// Call logout service with both tokens
	if err := u.userService.Logout(ctx, token, refreshToken); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, nil, "Logout successful")
}

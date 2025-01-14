package middleware

import (
	"go-fiber-api/internal/config"
	"go-fiber-api/internal/model"
	"go-fiber-api/internal/service"
	"go-fiber-api/pkg/utils"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	userService *service.UserService
	config      *config.Config
}

func NewAuthMiddleware(userService *service.UserService, config *config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		userService: userService,
		config:      config,
	}
}

// Protected validates JWT token and adds user to context
func (m *AuthMiddleware) Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.SendError(c, http.StatusUnauthorized, "Authorization header is required")
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			return utils.SendError(c, http.StatusUnauthorized, "Invalid token format")
		}

		token := bearerToken[1]
		auth := utils.NewAuthHandler(m.config.JWTSecretKey, m.config.JWTRefreshKey, m.config.JWTExpiresIn, m.config.JWTRefreshIn)
		claims, err := auth.ValidateToken(token)
		if err != nil {
			return utils.SendError(c, http.StatusUnauthorized, "Invalid token")
		}

		if err := m.userService.ValidateTokenWithRedis(c.Context(), token); err != nil {
			return utils.SendError(c, http.StatusUnauthorized, "Token is invalid or has been revoked")
		}

		user, err := m.userService.FindByID(c.Context(), claims.UserID)
		if err != nil {
			return utils.SendError(c, http.StatusUnauthorized, "User not found")
		}

		c.Locals("user", user)
		c.Locals("token", token)
		c.Locals("claims", claims)
		return c.Next()
	}
}

// RequireRoles checks if user has required roles
func (m *AuthMiddleware) RequireRoles(roles ...utils.Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(*model.User)
		if !ok {
			return utils.SendError(c, http.StatusUnauthorized, "User not found in context")
		}

		userRoles := make([]utils.Role, len(user.Roles))
		for i, r := range user.Roles {
			userRoles[i] = utils.Role(r)
		}

		if !utils.IsValidRole(userRoles, roles) {
			return utils.SendError(c, http.StatusForbidden, "Insufficient permissions")
		}

		return c.Next()
	}
}

// GetUserFromContext retrieves user from context
func GetUserFromContext(c *fiber.Ctx) (*model.User, bool) {
	user, ok := c.Locals("user").(*model.User)
	return user, ok
}

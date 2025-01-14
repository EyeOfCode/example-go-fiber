package routes

import (
	"go-fiber-api/internal/config"
	"go-fiber-api/internal/handlers"
	"go-fiber-api/pkg/middleware"
	"go-fiber-api/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"

	_ "go-fiber-api/docs"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type Application struct {
	App              *fiber.App
	UserHandler      *handlers.UserHandler
	ShopHandler      *handlers.ShopHandler
	CategoryHandler  *handlers.CategoryHandler
	FileStoreHandler *handlers.FileStoreHandler
	OtherHandler     *handlers.OtherHandler
	AuthMiddleware   *middleware.AuthMiddleware
	Config           *config.Config
}

func (app *Application) SetupRoutes() {
	// Swagger route
	app.App.Get("/swagger/*", fiberSwagger.WrapHandler)

	// API routes
	v1 := app.App.Group("/api/v1")
	// Rate limit (You can use route by route)
	v1.Use(middleware.RateLimit(100, time.Minute))

	// Public routes
	public := v1.Group("/")

	// Auth routes
	auth := v1.Group("/auth")
	auth.Post("/register", app.UserHandler.Register)
	auth.Post("/login", app.UserHandler.Login)

	// Other routes
	other := public.Group("/other")
	other.Use(middleware.RateLimit(20, time.Minute))
	other.Get("/example/gallery", app.OtherHandler.GetListImages)

	// Protected routes
	private := v1.Group("/")
	private.Use(app.AuthMiddleware.Protected())

	// User routes
	users := private.Group("/user")
	users.Get("/profile", app.UserHandler.GetProfile)

	// Auth routes
	user := private.Group("/auth")
	user.Get("/logout", app.UserHandler.Logout)
	user.Post("/refresh", app.UserHandler.RefreshToken)

	// Admin only routes
	adminGroup := private.Group("/admin")
	adminGroup.Use(app.AuthMiddleware.RequireRoles(utils.Role("admin")))
	adminGroup.Get("/users", app.UserHandler.UserList)
	adminGroup.Put("/user/:id", app.UserHandler.UpdateUser)
	adminGroup.Delete("/user/:id", app.UserHandler.DeleteUser)

	// Shop routes
	shops := private.Group("/shop")
	shops.Get("/list", app.ShopHandler.ShopList)
	shops.Post("/", app.ShopHandler.CreateShop)
	shops.Get("/:id", app.ShopHandler.GetShop)
	shops.Put("/:id", app.ShopHandler.UpdateShop)
	shops.Delete("/:id", app.ShopHandler.DeleteShop)

	// Category routes
	categories := private.Group("/category")
	categories.Get("/list", app.CategoryHandler.GetAll)
	categories.Post("/", app.CategoryHandler.Create)
	categories.Delete("/:id", app.CategoryHandler.DeleteCategory)

	// file routes
	files := private.Group("/file")
	files.Get("shop/:shop_id/download/:file_id", app.FileStoreHandler.Download)
}

package handlers

import (
	"context"
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

type ShopHandler struct {
    shopService *service.ShopService
}

func NewShopHandler(shopService *service.ShopService) *ShopHandler {
    return &ShopHandler{
        shopService: shopService,
    }
}

// @Summary List shops
// @Description Get paginated list of shops with optional filtering 
// @Tags shop
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int false "Page number" default(1) 
// @Param page_size query int false "Page size" default(10)
// @Param name query string false "Filter by name"
// @Success 200
// @Router /shop/list [get]
func (s *ShopHandler) ShopList(c *fiber.Ctx) error {
    page, pageSize := utils.PaginationParams(c)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    total, err := s.shopService.Count(ctx, bson.D{})
    if err != nil {
        return utils.SendError(c, http.StatusInternalServerError, "Failed to count shops: "+err.Error())
    }

    opts := options.Find().
        SetSkip(int64((page - 1) * pageSize)).
        SetLimit(int64(pageSize)).
        SetSort(bson.D{{Key: "created_at", Value: -1}})

    shops, err := s.shopService.FindAll(ctx, bson.M{}, opts)
    if err != nil {
        return utils.SendError(c, http.StatusInternalServerError, err.Error())
    }

    response := utils.CreatePagination(page, pageSize, total, shops)
    return utils.SendSuccess(c, http.StatusOK, response)
}

// @Summary Create Shop endpoint
// @Description Post the API's create shop
// @Tags shop
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body dto.ShopRequest true "Shop details"
// @Router /shop [post]
func (s *ShopHandler) CreateShop(c *fiber.Ctx) error {
    var req dto.ShopRequest

    if err := c.BodyParser(&req); err != nil {
        return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
    }

    if err := utils.ValidateStruct(&req); err != nil {
        return utils.SendValidationError(c, err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    user, ok := middleware.GetUserFromContext(c)
    if !ok {
        return utils.SendError(c, http.StatusUnauthorized, "Invalid session")
    }

    shop, err := s.shopService.Create(ctx, &req, user)
    if err != nil {
        return utils.SendError(c, http.StatusInternalServerError, err.Error())
    }

    res := &dto.UpdateShopResponse{
        ID:        shop.ID,
        Name:      shop.Name,
        Budget:    shop.Budget,
        CreatedAt: shop.CreatedAt,
        UpdatedAt: shop.UpdatedAt,
    }
    return utils.SendSuccess(c, http.StatusCreated, res)
}

// @Summary Get Shop endpoint
// @Description Get the API's get shop
// @Tags shop
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Shop ID"
// @Router /shop/{id} [get]
func (s *ShopHandler) GetShop(c *fiber.Ctx) error {
    id := c.Params("id")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return utils.SendError(c, http.StatusBadRequest, "Invalid ID format")
    }

    shop, err := s.shopService.FindByID(ctx, objID)
    if err != nil {
        return utils.SendError(c, http.StatusNotFound, "Failed to find shop")
    }

    return utils.SendSuccess(c, http.StatusOK, shop)
}

// @Summary Update Shop endpoint
// @Description Get the API's update shop
// @Tags shop
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Shop ID"
// @Param request body dto.ShopRequest true "Shop update details"
// @Router /shop/{id} [put]
func (s *ShopHandler) UpdateShop(c *fiber.Ctx) error {
    var req dto.ShopRequest

    if err := c.BodyParser(&req); err != nil {
        return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
    }

    if err := utils.ValidateStruct(&req); err != nil {
        return utils.SendValidationError(c, err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    user, ok := middleware.GetUserFromContext(c)
    if !ok {
        return utils.SendError(c, http.StatusUnauthorized, "Invalid session")
    }

    id := c.Params("id")
    shopId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return utils.SendError(c, http.StatusBadRequest, "Invalid ID format")
    }

    shop, err := s.shopService.FindByID(ctx, shopId)
    if err != nil || shop == nil {
        return utils.SendError(c, http.StatusNotFound, "Failed to find shop")
    }

    if shop.CreatedBy != user.ID {
        return utils.SendError(c, http.StatusUnauthorized, "Unauthorized")
    }

    shop, err = s.shopService.Update(ctx, shopId, &req)
    if err != nil {
        return utils.SendError(c, http.StatusInternalServerError, err.Error())
    }

    res := &dto.UpdateShopResponse{
        ID:        shop.ID,
        Name:      shop.Name,
        Budget:    shop.Budget,
        CreatedAt: shop.CreatedAt,
        UpdatedAt: shop.UpdatedAt,
    }
    return utils.SendSuccess(c, http.StatusOK, res)
}

// @Summary Delete Shop endpoint
// @Description Get the API's delete shop
// @Tags shop
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Shop ID"
// @Router /shop/{id} [delete]
func (s *ShopHandler) DeleteShop(c *fiber.Ctx) error {
    id := c.Params("id")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    shopId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return utils.SendError(c, http.StatusBadRequest, "Invalid ID format")
    }

    shop, err := s.shopService.FindByID(ctx, shopId)
    if err != nil || shop == nil {
        return utils.SendError(c, http.StatusNotFound, "Failed to find shop")
    }

    user, ok := middleware.GetUserFromContext(c)
    if !ok {
        return utils.SendError(c, http.StatusUnauthorized, "Invalid session")
    }

    if shop.CreatedBy != user.ID {
        return utils.SendError(c, http.StatusUnauthorized, "Unauthorized")
    }

    err = s.shopService.Delete(ctx, shopId)
    if err != nil {
        return utils.SendError(c, http.StatusInternalServerError, "test")
    }

    return utils.SendSuccess(c, http.StatusOK, nil, "Shop deleted successfully")
}
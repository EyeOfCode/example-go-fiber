package handlers

import (
	"context"
	"go-fiber-api/internal/service"
	"go-fiber-api/pkg/utils"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileStoreHandler struct {
	fileStoreService *service.FileStoreService
	shopService      *service.ShopService
}

func NewFileStoreHandler(fileStoreService *service.FileStoreService, shopService *service.ShopService) *FileStoreHandler {
	return &FileStoreHandler{
		fileStoreService: fileStoreService,
		shopService:      shopService,
	}
}

// @Summary Download File Store endpoint
// @Description Get the API's download file store
// @Tags file-store
// @Accept json
// @Produce octet-stream
// @Security Bearer
// @Param shop_id path string true "Shop ID"
// @Param file_id path string true "File Store ID"
// @Router /file/shop/{shop_id}/download/{file_id} [get]
func (f *FileStoreHandler) Download(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	shopId := c.Params("shop_id")

	objID, err := primitive.ObjectIDFromHex(shopId)
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid ID format")
	}

	shop, err := f.shopService.FindByID(ctx, objID)
	if err != nil || shop == nil {
		return utils.SendError(c, http.StatusNotFound, "Failed to find file store")
	}

	id := c.Params("file_id")

	objID, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid ID format")
	}

	fileStore, err := f.fileStoreService.FindOne(ctx, bson.M{"_id": objID})
	if err != nil || fileStore == nil {
		return utils.SendError(c, http.StatusNotFound, "Failed to find file store")
	}

	fullPath := filepath.Join(fileStore.BasePath, fileStore.Name)
	return c.Download(fullPath)
}

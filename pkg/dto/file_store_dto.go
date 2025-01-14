package dto

import (
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileRequest struct {
	Files []multipart.FileHeader `form:"files" binding:"omitempty"`
}

type FileStoreRequest struct {
	Files  []multipart.FileHeader `json:"files" binding:"required"`
	ShopID primitive.ObjectID     `json:"shop_id" binding:"required"`
}

type FileStoreUploadResponse struct {
	Name      string `json:"name"`
	BasePath  string `json:"base_path"`
	Extension string `json:"extension"`
}

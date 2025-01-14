package dto

import (
	"go-fiber-api/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShopRequest struct {
	Name   string  `json:"name" form:"name" binding:"required,min=3,max=30"`
	Budget float64 `json:"budget" form:"budget" binding:"required"`
}

type UpdateShopRequest struct {
	Name   string  `json:"name" form:"name" binding:"omitempty,min=3,max=30"`
	Budget float64 `json:"budget" form:"budget" binding:"omitempty"`
}

type UpdateShopResponse struct {
	ID        primitive.ObjectID `json:"_id"`
	Name      string             `json:"name"`
	Budget    float64            `json:"budget"`
	Files     []*model.FileStore `json:"files"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

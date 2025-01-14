package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileStore struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	BasePath  string             `json:"base_path" bson:"base_path"`
	Extension string             `json:"extension" bson:"extension"`
	ShopID    primitive.ObjectID `json:"shop_id" bson:"shop_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

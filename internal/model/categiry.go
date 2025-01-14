package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	ShopID    primitive.ObjectID `json:"shop_id" bson:"shop_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Price       float64            `json:"price" bson:"price"`
	ImageID     primitive.ObjectID `json:"image_id" bson:"image_id,omitempty"`
	CategoryID  primitive.ObjectID `json:"category_id" bson:"category_id,omitempty"`
	ShopID      primitive.ObjectID `json:"shop_id" bson:"shop_id"`
}

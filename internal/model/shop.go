package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shop struct {
	ID         primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Name       string              `bson:"name" json:"name"`
	Budget     float64             `bson:"budget,default=0" json:"budget"`
	CreatedBy  primitive.ObjectID  `bson:"created_by" json:"created_by"`
	User       *UserResponseOnShop `bson:"user,omitempty" json:"user"`
	Categories []*Category         `bson:"categories,omitempty" json:"categories,omitempty"`
	Files      []*FileStore        `bson:"files,omitempty" json:"files,omitempty"`
	CreatedAt  time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time           `bson:"updated_at" json:"updated_at"`
}

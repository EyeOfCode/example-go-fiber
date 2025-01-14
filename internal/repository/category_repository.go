package repository

import (
	"context"
	"go-fiber-api/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *model.Category) (*model.Category, error)
	Get(ctx context.Context, id primitive.ObjectID) (*model.Category, error)
	List(ctx context.Context) ([]model.Category, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type categoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository(db *mongo.Database) CategoryRepository {
	return &categoryRepository{
		collection: db.Collection("categories"),
	}
}

func (r *categoryRepository) Create(ctx context.Context, category *model.Category) (*model.Category, error) {
	category.ID = primitive.NewObjectID()
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) Get(ctx context.Context, id primitive.ObjectID) (*model.Category, error) {
	var category model.Category
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) List(ctx context.Context) ([]model.Category, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []model.Category
	if err := cursor.All(ctx, &categories); err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

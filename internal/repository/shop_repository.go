package repository

import (
	"context"
	"go-fiber-api/internal/model"
	"go-fiber-api/pkg/dto"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShopRepository interface {
	Create(ctx context.Context, shop *model.Shop) (*model.Shop, error)
	FindOne(ctx context.Context, query bson.M) (*model.Shop, error)
	FindAll(ctx context.Context, query bson.M, opts *options.FindOptions) ([]model.Shop, error)
	Count(ctx context.Context, query bson.D) (int64, error)
	UpdateByID(ctx context.Context, id primitive.ObjectID, payload *dto.UpdateShopRequest) (*model.Shop, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type shopRepository struct {
	collection *mongo.Collection
}

func NewShopRepository(db *mongo.Database) ShopRepository {
	return &shopRepository{
		collection: db.Collection("shops"),
	}
}

func (r *shopRepository) Create(ctx context.Context, shop *model.Shop) (*model.Shop, error) {
	shop.ID = primitive.NewObjectID()
	shop.CreatedAt = time.Now()
	shop.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, shop)
	if err != nil {
		return nil, err
	}
	return shop, nil
}

func (r *shopRepository) FindOne(ctx context.Context, query bson.M) (*model.Shop, error) {
	var shop model.Shop
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: query}},
		{{Key: "$sort", Value: bson.D{{Key: "created_at", Value: -1}}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "created_by",
			"foreignField": "_id",
			"as":           "user",
		}}},
		{{Key: "$unwind", Value: "$user"}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "categories",
			"localField":   "_id",
			"foreignField": "shop_id",
			"as":           "categories",
		}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "file_stores",
			"localField":   "_id",
			"foreignField": "shop_id",
			"as":           "files",
		}}},
	}
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if !cursor.Next(ctx) {
		return nil, nil
	}

	err = cursor.Decode(&shop)
	if err != nil {
		return nil, err
	}

	return &shop, nil
}

func (r *shopRepository) FindAll(ctx context.Context, query bson.M, opts *options.FindOptions) ([]model.Shop, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: query}},
		{{Key: "$sort", Value: bson.D{{Key: "created_at", Value: -1}}}},
		{{Key: "$skip", Value: opts.Skip}},
		{{Key: "$limit", Value: opts.Limit}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "users",
			"localField":   "created_by",
			"foreignField": "_id",
			"as":           "user",
		}}},
		{{Key: "$unwind", Value: "$user"}},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var shops []model.Shop
	if err := cursor.All(ctx, &shops); err != nil {
		return nil, err
	}
	return shops, nil
}

func (r *shopRepository) Count(ctx context.Context, query bson.D) (int64, error) {
	return r.collection.CountDocuments(ctx, query)
}

func (r *shopRepository) UpdateByID(ctx context.Context, id primitive.ObjectID, payload *dto.UpdateShopRequest) (*model.Shop, error) {
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedShop model.Shop
	err := r.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": id},
		bson.M{
			"$set": payload,
			"$currentDate": bson.M{
				"updated_at": true,
			},
		},
		opts,
	).Decode(&updatedShop)
	if err != nil {
		return nil, err
	}
	return &updatedShop, nil
}

func (r *shopRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

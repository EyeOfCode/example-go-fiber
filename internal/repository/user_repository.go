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

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	UpdateByID(ctx context.Context, id primitive.ObjectID, payload *dto.UpdateUserRequest) (*model.User, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
	FindOne(ctx context.Context, query bson.M) (*model.User, error)
	FindAll(ctx context.Context, query bson.D, opts *options.FindOptions) ([]model.User, error)
	Count(ctx context.Context, query bson.D) (int64, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *userRepository) FindOne(ctx context.Context, query bson.M) (*model.User, error) {
	var user model.User
	err := r.collection.FindOne(ctx, query).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateByID(ctx context.Context, id primitive.ObjectID, payload *dto.UpdateUserRequest) (*model.User, error) {
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedUser model.User
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
	).Decode(&updatedUser)
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func (r *userRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *userRepository) FindAll(ctx context.Context, query bson.D, opts *options.FindOptions) ([]model.User, error) {
	cursor, err := r.collection.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []model.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Count(ctx context.Context, query bson.D) (int64, error) {
	return r.collection.CountDocuments(ctx, query)
}

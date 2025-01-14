package repository

import (
	"context"
	"go-fiber-api/internal/model"
	"os"
	"path/filepath"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FileStoreRepository interface {
	FindAll(ctx context.Context, query bson.M) ([]model.FileStore, error)
	FindById(ctx context.Context, id primitive.ObjectID) (*model.FileStore, error)
	FindOne(ctx context.Context, query bson.M) (*model.FileStore, error)
	Create(ctx context.Context, fileStore []*model.FileStore) ([]*model.FileStore, error)
	Delete(ctx context.Context, fileStore *model.FileStore, id primitive.ObjectID) error
}

type fileStoreRepository struct {
	collection *mongo.Collection
}

func NewFileStoreRepository(db *mongo.Database) FileStoreRepository {
	return &fileStoreRepository{
		collection: db.Collection("file_stores"),
	}
}

func (r *fileStoreRepository) FindAll(ctx context.Context, query bson.M) ([]model.FileStore, error) {
	var fileStores []model.FileStore
	cursor, err := r.collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &fileStores); err != nil {
		return nil, err
	}
	return fileStores, nil
}

func (r *fileStoreRepository) FindOne(ctx context.Context, query bson.M) (*model.FileStore, error) {
	var result model.FileStore
	err := r.collection.FindOne(ctx, query).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *fileStoreRepository) FindById(ctx context.Context, id primitive.ObjectID) (*model.FileStore, error) {
	var result model.FileStore
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *fileStoreRepository) Create(ctx context.Context, fileStore []*model.FileStore) ([]*model.FileStore, error) {
	if len(fileStore) == 0 {
		return nil, nil
	}
	now := time.Now()
	documents := make([]interface{}, len(fileStore))
	for i, fs := range fileStore {
		fs.CreatedAt = now
		fs.UpdatedAt = now
		documents[i] = fs
	}
	result, err := r.collection.InsertMany(ctx, documents)
	if err != nil {
		return nil, err
	}
	for i, insertedID := range result.InsertedIDs {
		if oid, ok := insertedID.(primitive.ObjectID); ok {
			fileStore[i].ID = oid
		}
	}
	return fileStore, nil
}

func (r *fileStoreRepository) Delete(ctx context.Context, fileStore *model.FileStore, id primitive.ObjectID) error {
	filePath := filepath.Join(fileStore.BasePath, fileStore.Name)
	if err := os.Remove(filePath); err != nil {
		return err
	}
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

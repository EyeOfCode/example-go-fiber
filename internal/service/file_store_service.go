package service

import (
	"context"
	"go-fiber-api/internal/model"
	"go-fiber-api/internal/repository"
	"go-fiber-api/pkg/dto"
	"go-fiber-api/pkg/utils"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileStoreService struct {
	fileStoreRepo repository.FileStoreRepository
}

func NewFileStoreService(fileStoreRepo repository.FileStoreRepository) *FileStoreService {
	return &FileStoreService{
		fileStoreRepo: fileStoreRepo,
	}
}
func (s *FileStoreService) Uploads(ctx context.Context, payload *dto.FileStoreRequest, shop *model.Shop) ([]*model.FileStore, error) {
	var files []*multipart.FileHeader
	for i := range payload.Files {
		files = append(files, &payload.Files[i])
	}
	resUpload, err := utils.Upload(files)
	if err != nil {
		return nil, err
	}

	var fileStore []*model.FileStore
	for i := range resUpload {
		fileStore = append(fileStore, &model.FileStore{
			ID:        primitive.NewObjectID(),
			Name:      resUpload[i].Name,
			BasePath:  resUpload[i].BasePath,
			Extension: resUpload[i].Extension,
			ShopID:    shop.ID,
		})
	}

	createdFileStore, err := s.fileStoreRepo.Create(ctx, fileStore)
	if err != nil {
		return nil, err
	}
	return createdFileStore, nil
}

func (s *FileStoreService) Delete(ctx context.Context, id primitive.ObjectID) error {
	fileStore, err := s.fileStoreRepo.FindById(ctx, id)
	if err != nil {
		return err
	}
	return s.fileStoreRepo.Delete(ctx, fileStore, id)
}

func (s *FileStoreService) FindAll(ctx context.Context, query bson.M) ([]model.FileStore, error) {
	return s.fileStoreRepo.FindAll(ctx, query)
}

func (s *FileStoreService) FindOne(ctx context.Context, query bson.M) (*model.FileStore, error) {
	return s.fileStoreRepo.FindOne(ctx, query)
}

package service

import (
	"context"
	"go-fiber-api/internal/model"
	"go-fiber-api/internal/repository"
	"go-fiber-api/pkg/dto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *CategoryService) Create(ctx context.Context, payload *dto.CategoryRequest, shop *model.Shop) (*model.Category, error) {
	category := &model.Category{
		Name:   payload.Name,
		ShopID: shop.ID,
	}

	createdCategory, err := s.categoryRepo.Create(ctx, category)
	if err != nil {
		return nil, err
	}
	return createdCategory, nil
}

func (s *CategoryService) FindAll(ctx context.Context) ([]model.Category, error) {
	return s.categoryRepo.List(ctx)
}

func (s *CategoryService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.categoryRepo.Delete(ctx, id)
}

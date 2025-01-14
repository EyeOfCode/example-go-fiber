package service

import (
	"context"
	"go-fiber-api/internal/model"
	"go-fiber-api/internal/repository"
	"go-fiber-api/pkg/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShopService struct {
	shopRepo repository.ShopRepository
}

func NewShopService(shopRepo repository.ShopRepository) *ShopService {
	return &ShopService{
		shopRepo: shopRepo,
	}
}
func (s *ShopService) FindAll(ctx context.Context, query bson.M, opts *options.FindOptions) ([]model.Shop, error) {
	return s.shopRepo.FindAll(ctx, query, opts)
}

func (s *ShopService) Count(ctx context.Context, query bson.D) (int64, error) {
	return s.shopRepo.Count(ctx, query)
}

func (s *ShopService) Create(ctx context.Context, payload *dto.ShopRequest, user *model.User) (*model.Shop, error) {
	shop := &model.Shop{
		Name:      payload.Name,
		Budget:    payload.Budget,
		CreatedBy: user.ID,
	}
	createdShop, err := s.shopRepo.Create(ctx, shop)
	if err != nil {
		return nil, err
	}
	return createdShop, nil
}

func (s *ShopService) FindByID(ctx context.Context, id primitive.ObjectID) (*model.Shop, error) {
	return s.shopRepo.FindOne(ctx, bson.M{"_id": id})
}

func (s *ShopService) Update(ctx context.Context, id primitive.ObjectID, payload *dto.ShopRequest) (*model.Shop, error) {
	shop := &dto.UpdateShopRequest{
		Name:   payload.Name,
		Budget: payload.Budget,
	}

	updatedShop, err := s.shopRepo.UpdateByID(ctx, id, shop)
	if err != nil {
		return nil, err
	}
	return updatedShop, nil
}

func (s *ShopService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.shopRepo.Delete(ctx, id)
}

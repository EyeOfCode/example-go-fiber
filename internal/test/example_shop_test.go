package test

import (
	"context"
	"go-fiber-api/internal/model"
	"go-fiber-api/internal/service"
	"go-fiber-api/pkg/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockShopRepository struct {
	mock.Mock
}

func (m *MockShopRepository) Create(ctx context.Context, shop *model.Shop) (*model.Shop, error) {
	args := m.Called(ctx, shop)
	return args.Get(0).(*model.Shop), args.Error(1)
}

func (m *MockShopRepository) FindAll(ctx context.Context, query bson.M, opts *options.FindOptions) ([]model.Shop, error) {
	return nil, nil
}

func (m *MockShopRepository) Count(ctx context.Context, query bson.D) (int64, error) {
	return 0, nil
}

func (m *MockShopRepository) FindOne(ctx context.Context, query bson.M) (*model.Shop, error) {
	return nil, nil
}

func (m *MockShopRepository) UpdateByID(ctx context.Context, id primitive.ObjectID, update *dto.UpdateShopRequest) (*model.Shop, error) {
	return nil, nil
}

func (m *MockShopRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	return nil
}

func TestShopService_Create(t *testing.T) {
	mockRepo := &MockShopRepository{}
	shopService := service.NewShopService(mockRepo)

	ctx := context.Background()
	userID := primitive.NewObjectID()
	payload := &dto.ShopRequest{
		Name:   "Test Shop",
		Budget: 1000,
	}
	user := &model.User{ID: userID}

	expectedShop := &model.Shop{
		Name:      payload.Name,
		Budget:    payload.Budget,
		CreatedBy: userID,
	}

	mockRepo.On("Create", ctx, mock.MatchedBy(func(s *model.Shop) bool {
		return s.Name == expectedShop.Name &&
			s.Budget == expectedShop.Budget &&
			s.CreatedBy == expectedShop.CreatedBy
	})).Return(expectedShop, nil)

	result, err := shopService.Create(ctx, payload, user)

	assert.NoError(t, err)
	assert.Equal(t, expectedShop, result)
	mockRepo.AssertExpectations(t)
}

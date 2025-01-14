package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go-fiber-api/internal/config"
	"go-fiber-api/internal/repository"
	"go-fiber-api/pkg/dto"
)

type ArtworkApiService struct {
	httpServiceRepo repository.HttpServiceRepository
	config          *config.Config
}

func NewArtworkApiService(httpServiceRepo repository.HttpServiceRepository, config *config.Config) *ArtworkApiService {
	return &ArtworkApiService{
		httpServiceRepo: httpServiceRepo,
		config:          config,
	}
}

func (s *ArtworkApiService) GetListImages(ctx context.Context) (*dto.ArtworkResponse, error) {
	url := s.config.ArtworkApiURL
	rawResponse, err := s.httpServiceRepo.Get(ctx, url)
	if err != nil {
		return nil, err
	}

	// Marshal the raw response back to JSON
	jsonBytes, err := json.Marshal(rawResponse)
	if err != nil {
		return nil, err
	}

	// Unmarshal into the expected type
	var response dto.ArtworkResponse
	if err := json.Unmarshal(jsonBytes, &response); err != nil {
		return nil, err
	}

	if len(response.Data) == 0 {
		return nil, fmt.Errorf("no artworks found")
	}

	return &response, nil
}

package repository

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type HttpServiceRepository interface {
	Get(ctx context.Context, url string) (interface{}, error)
}

type httpServiceRepository struct {
	client *http.Client
}

func NewHttpServiceRepository() HttpServiceRepository {
	return &httpServiceRepository{
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (s *httpServiceRepository) Get(ctx context.Context, url string) (any, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

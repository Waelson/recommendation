package service

import (
	"encoding/json"
	"fmt"
	"github.com/Waelson/recommendation/ecommerce-api/internal/model"
	"net/http"
	"net/url"
)

type RecommendationService struct {
	BaseURL string
}

// NewRecommendationService creates a new instance of RecommendationService
func NewRecommendationService(baseURL string) RecommendationService {
	return RecommendationService{BaseURL: baseURL}
}

// GetRecommendations fetches product recommendations from the endpoint
func (s *RecommendationService) GetRecommendations(productID string, n int) (*model.RecommendationResponse, error) {
	// Construct the request URL with query parameters
	path := fmt.Sprintf("%s/recommend", s.BaseURL)
	reqURL, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %v", err)
	}

	query := reqURL.Query()
	query.Set("product_id", productID)
	query.Set("n", fmt.Sprintf("%d", n))
	reqURL.RawQuery = query.Encode()

	// Make the HTTP GET request
	resp, err := http.Get(reqURL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK response: %s", resp.Status)
	}

	// Decode the response body into the RecommendationResponse struct
	var recommendations model.RecommendationResponse
	err = json.NewDecoder(resp.Body).Decode(&recommendations)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %v", err)
	}

	return &recommendations, nil
}

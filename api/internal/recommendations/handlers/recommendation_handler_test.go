package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"stonks-api/internal/recommendations/handlers"
	"stonks-api/internal/recommendations/mocks"
	"stonks-api/internal/recommendations/services"
	"stonks-api/internal/stocks/models"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestGetRecommendations(t *testing.T) {
	// Successful recommendations
	t.Run("successful recommendations", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/stonks-api/recommendations", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		recommendations := []services.StockRecommendation{
			{
				Stock: models.Stock{
					Ticker:  "AAPL",
					Company: "Apple Inc.",
				},
				Score:  4.5,
				Reason: "Stock was recently upgraded, Target price increased",
			},
			{
				Stock: models.Stock{
					Ticker:  "MSFT",
					Company: "Microsoft Corp",
				},
				Score:  3.2,
				Reason: "Target price increased",
			},
		}

		mockService := &mocks.MockRecommendationService{
			GetRecommendationsFn: func() ([]services.StockRecommendation, error) {
				return recommendations, nil
			},
		}

		h := handlers.NewRecommendationHandler(mockService)

		// Act
		err := h.GetRecommendations(c)

		// Assert
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, rec.Code)
		}

		var response []services.StockRecommendation
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		if err != nil {
			t.Errorf("Error unmarshaling response: %v", err)
		}

		if len(response) != len(recommendations) {
			t.Errorf("Expected %d recommendations but got %d", len(recommendations), len(response))
		}

		if response[0].Stock.Ticker != "AAPL" || response[1].Stock.Ticker != "MSFT" {
			t.Errorf("Response did not match expected recommendations")
		}
	})

	// No recommendations available
	t.Run("no recommendations", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/stonks-api/recommendations", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockService := &mocks.MockRecommendationService{
			GetRecommendationsFn: func() ([]services.StockRecommendation, error) {
				return []services.StockRecommendation{}, nil
			},
		}

		h := handlers.NewRecommendationHandler(mockService)

		// Act
		err := h.GetRecommendations(c)

		// Assert
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, rec.Code)
		}

		var response map[string]string
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		if err != nil {
			t.Errorf("Error unmarshaling response: %v", err)
		}

		message, exists := response["message"]
		if !exists || message != "No recommendations available at this time" {
			t.Errorf("Expected message about no recommendations but got: %v", response)
		}
	})

	// Service error
	t.Run("service error", func(t *testing.T) {
		// Setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/stonks-api/recommendations", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockService := &mocks.MockRecommendationService{
			GetRecommendationsFn: func() ([]services.StockRecommendation, error) {
				return nil, errors.New("service error")
			},
		}

		h := handlers.NewRecommendationHandler(mockService)

		// Act
		err := h.GetRecommendations(c)

		// Assert
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, rec.Code)
		}

		var response map[string]string
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		if err != nil {
			t.Errorf("Error unmarshaling response: %v", err)
		}

		errMsg, exists := response["error"]
		if !exists || errMsg != "Failed to get recommendations: service error" {
			t.Errorf("Expected error message but got: %v", response)
		}
	})
}

func TestRegisterRoutes(t *testing.T) {
	t.Run("register routes", func(t *testing.T) {
		// Setup
		e := echo.New()
		group := e.Group("/api/v1/stonks-api")

		mockService := &mocks.MockRecommendationService{}
		h := handlers.NewRecommendationHandler(mockService)

		// no error means routes were registered successfully
		h.RegisterRoutes(group)

		// Verify the route exists and returns the expected handler
		routes := e.Routes()
		found := false
		for _, route := range routes {
			if route.Path == "/api/v1/stonks-api/recommendations" && route.Method == http.MethodGet {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Recommendation route was not registered")
		}
	})
}

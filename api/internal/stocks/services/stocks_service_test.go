package services_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"stonks-api/internal/stocks/models"
	"stonks-api/internal/stocks/services"
	"testing"
	"time"
)

// MockRepository implements the Repository interface for testing
type MockRepository struct {
	SaveStocksFn        func(stocks []models.Stock) error
	GetAllStocksFn      func(params models.PaginationParams) (models.PaginatedStocks, error)
	GetStocksByTickerFn func(ticker string) ([]models.Stock, error)
	GetRecentStocksFn   func(limit int) ([]models.Stock, error)
}

func (m *MockRepository) SaveStocks(stocks []models.Stock) error {
	if m.SaveStocksFn != nil {
		return m.SaveStocksFn(stocks)
	}
	return nil
}

func (m *MockRepository) GetAllStocks(params models.PaginationParams) (models.PaginatedStocks, error) {
	if m.GetAllStocksFn != nil {
		return m.GetAllStocksFn(params)
	}
	return models.PaginatedStocks{}, nil
}

func (m *MockRepository) GetStocksByTicker(ticker string) ([]models.Stock, error) {
	if m.GetStocksByTickerFn != nil {
		return m.GetStocksByTickerFn(ticker)
	}
	return []models.Stock{}, nil
}

func (m *MockRepository) GetRecentStocks(limit int) ([]models.Stock, error) {
	if m.GetRecentStocksFn != nil {
		return m.GetRecentStocksFn(limit)
	}
	return []models.Stock{}, nil
}

// MockHTTPClient implements http client for testing
type MockHTTPClient struct {
	DoFn func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFn != nil {
		return m.DoFn(req)
	}
	return nil, errors.New("not implemented")
}

func TestFetchStocks(t *testing.T) {
	// Successful fetch
	t.Run("successful fetch", func(t *testing.T) {
		mockResp := services.StockResponse{
			Items: []services.StockItem{
				{
					Ticker:     "AAPL",
					Company:    "Apple Inc.",
					TargetFrom: "$150.00",
					TargetTo:   "$200.00",
				},
			},
		}

		// Create mock response
		rec := httptest.NewRecorder()
		json.NewEncoder(rec).Encode(mockResp)

		mockClient := &MockHTTPClient{
			DoFn: func(req *http.Request) (*http.Response, error) {
				return rec.Result(), nil
			},
		}

		mockRepo := &MockRepository{}
		service := services.NewStockService(mockRepo)
		service.SetHTTPClient(mockClient)

		resp, err := service.FetchStocks("")

		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}

		if resp == nil || len(resp.Items) != 1 {
			t.Errorf("Expected 1 item in response")
		}
	})

	// API error
	t.Run("api error", func(t *testing.T) {
		mockClient := &MockHTTPClient{
			DoFn: func(req *http.Request) (*http.Response, error) {
				return nil, errors.New("connection error")
			},
		}

		mockRepo := &MockRepository{}
		service := services.NewStockService(mockRepo)
		service.SetHTTPClient(mockClient)

		_, err := service.FetchStocks("")

		if err == nil {
			t.Errorf("Expected error but got nil")
		}
	})
}

func TestSyncStocks(t *testing.T) {
	// Successful sync
	t.Run("successful sync", func(t *testing.T) {
		// Create mock response with 2 items
		mockResp := services.StockResponse{
			Items: []services.StockItem{
				{Ticker: "AAPL", TargetFrom: "$150.00", Time: time.Now()},
				{Ticker: "MSFT", TargetFrom: "$200.00", Time: time.Now()},
			},
		}

		rec := httptest.NewRecorder()
		json.NewEncoder(rec).Encode(mockResp)

		mockClient := &MockHTTPClient{
			DoFn: func(req *http.Request) (*http.Response, error) {
				return rec.Result(), nil
			},
		}

		mockRepo := &MockRepository{
			SaveStocksFn: func(stocks []models.Stock) error {
				return nil
			},
		}

		service := services.NewStockService(mockRepo)
		service.SetHTTPClient(mockClient)

		count, err := service.SyncStocks()

		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}

		if count != 2 {
			t.Errorf("Expected count 2 but got %d", count)
		}
	})

	// Database error
	t.Run("database error", func(t *testing.T) {
		mockResp := services.StockResponse{
			Items: []services.StockItem{
				{Ticker: "AAPL", TargetFrom: "$150.00", Time: time.Now()},
			},
		}

		rec := httptest.NewRecorder()
		json.NewEncoder(rec).Encode(mockResp)

		mockClient := &MockHTTPClient{
			DoFn: func(req *http.Request) (*http.Response, error) {
				return rec.Result(), nil
			},
		}

		mockRepo := &MockRepository{
			SaveStocksFn: func(stocks []models.Stock) error {
				return errors.New("database error")
			},
		}

		service := services.NewStockService(mockRepo)
		service.SetHTTPClient(mockClient)

		_, err := service.SyncStocks()

		if err == nil {
			t.Errorf("Expected error but got nil")
		}
	})
}

func TestGetStocksByTicker(t *testing.T) {
	// Stock found
	t.Run("stock found", func(t *testing.T) {
		expectedStocks := []models.Stock{
			{
				Ticker:  "AAPL",
				Company: "Apple Inc.",
			},
		}

		mockRepo := &MockRepository{
			GetStocksByTickerFn: func(ticker string) ([]models.Stock, error) {
				return expectedStocks, nil
			},
		}

		service := services.NewStockService(mockRepo)

		stocks, err := service.GetStocksByTicker("AAPL")

		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
			return
		}

		if len(stocks) != 1 {
			t.Errorf("Expected 1 stock but got %d", len(stocks))
			return
		}

		if stocks[0].Ticker != "AAPL" {
			t.Errorf("Expected ticker AAPL but got %s", stocks[0].Ticker)
		}
	})

	// Repository error
	t.Run("repository error", func(t *testing.T) {
		mockRepo := &MockRepository{
			GetStocksByTickerFn: func(ticker string) ([]models.Stock, error) {
				return nil, errors.New("repository error")
			},
		}

		service := services.NewStockService(mockRepo)

		_, err := service.GetStocksByTicker("AAPL")

		if err == nil {
			t.Errorf("Expected error but got nil")
		}
	})
}

// File: internal/recommendations/mocks/mocks.go
package mocks

import (
	"stonks-api/internal/recommendations/services"
	"stonks-api/internal/stocks/models"
)

// MockStockRepository implements the interfaces.StockRepository interface for testing
type MockStockRepository struct {
	GetRecentStocksFn   func(limit int) ([]models.Stock, error)
	SaveStocksFn        func(stocks []models.Stock) error
	GetAllStocksFn      func(params models.PaginationParams) (models.PaginatedStocks, error)
	GetStocksByTickerFn func(ticker string) ([]models.Stock, error)
}

// GetRecentStocks implements the required method
func (m *MockStockRepository) GetRecentStocks(limit int) ([]models.Stock, error) {
	if m.GetRecentStocksFn != nil {
		return m.GetRecentStocksFn(limit)
	}
	return []models.Stock{}, nil
}

// SaveStocks implements the required method
func (m *MockStockRepository) SaveStocks(stocks []models.Stock) error {
	if m.SaveStocksFn != nil {
		return m.SaveStocksFn(stocks)
	}
	return nil
}

// GetAllStocks implements the required method
func (m *MockStockRepository) GetAllStocks(params models.PaginationParams) (models.PaginatedStocks, error) {
	if m.GetAllStocksFn != nil {
		return m.GetAllStocksFn(params)
	}
	return models.PaginatedStocks{}, nil
}

// GetStocksByTicker implements the required method
func (m *MockStockRepository) GetStocksByTicker(ticker string) ([]models.Stock, error) {
	if m.GetStocksByTickerFn != nil {
		return m.GetStocksByTickerFn(ticker)
	}
	return []models.Stock{}, nil
}

// MockRecommendationService implements the RecommendationServiceInterface for testing
type MockRecommendationService struct {
	GetRecommendationsFn func() ([]services.StockRecommendation, error)
}

// GetRecommendations implements the required method
func (m *MockRecommendationService) GetRecommendations() ([]services.StockRecommendation, error) {
	if m.GetRecommendationsFn != nil {
		return m.GetRecommendationsFn()
	}
	return []services.StockRecommendation{}, nil
}

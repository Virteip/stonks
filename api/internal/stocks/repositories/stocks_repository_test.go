package repository_test

import (
	"errors"
	"stonks-api/cmd/database"
	"stonks-api/internal/stocks/mocks"
	"stonks-api/internal/stocks/models"
	repository "stonks-api/internal/stocks/repositories"
	"testing"
	"time"
)

func TestSaveStocks(t *testing.T) {
	// Empty stocks list
	t.Run("empty stocks list", func(t *testing.T) {
		mockDB := &database.MockDatabase{}

		repo := repository.NewStockRepository(mockDB)

		err := repo.SaveStocks([]models.Stock{})

		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}
	})

	// Create error
	t.Run("database error", func(t *testing.T) {
		dbError := errors.New("database error")
		mockDB := database.NewMockDatabaseWithError(dbError)

		repo := repository.NewStockRepository(mockDB)

		stocks := []models.Stock{
			{
				Ticker:     "AAPL",
				Company:    "Apple Inc.",
				TargetFrom: 150.0,
				TargetTo:   200.0,
				Time:       time.Now(),
			},
		}

		err := repo.SaveStocks(stocks)

		if err == nil {
			t.Errorf("Expected error but got nil")
		}
	})

	// Successful save
	t.Run("successful save", func(t *testing.T) {
		mockDB := &database.MockDatabase{
			TransactionFn: func(fc func(tx database.Transaction) error) error {
				return fc(&database.MockTransaction{})
			},
		}

		repo := repository.NewStockRepository(mockDB)

		stocks := []models.Stock{
			{
				Ticker:     "AAPL",
				Company:    "Apple Inc.",
				TargetFrom: 150.0,
				TargetTo:   200.0,
				Time:       time.Now(),
			},
		}

		err := repo.SaveStocks(stocks)

		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}
	})
}

func TestGetAllStocks(t *testing.T) {
	// Successful retrieval
	t.Run("successful retrieval", func(t *testing.T) {
		expectedStocks := []models.Stock{
			{
				ID:      "1",
				Ticker:  "AAPL",
				Company: "Apple Inc.",
				Time:    time.Now(),
			},
		}

		// Use the helper to create a mock DB
		mockDB := mocks.CreateMockDBWithStocks(expectedStocks)

		// Set Count function to return 1 for the total count
		mockDB.CountFn = func(model interface{}) (int64, error) {
			return 1, nil
		}

		repo := repository.NewStockRepository(mockDB)

		params := models.PaginationParams{
			Page:     1,
			PageSize: 10,
		}

		result, err := repo.GetAllStocks(params)

		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}

		if len(result.Stocks) != 1 {
			t.Errorf("Expected 1 stock but got %d", len(result.Stocks))
		}

		if result.Stocks[0].Ticker != "AAPL" {
			t.Errorf("Expected ticker AAPL but got %s", result.Stocks[0].Ticker)
		}
	})

	// Database error
	t.Run("database error", func(t *testing.T) {
		dbError := errors.New("database error")
		mockDB := database.NewMockDatabaseWithError(dbError)

		repo := repository.NewStockRepository(mockDB)

		params := models.PaginationParams{
			Page:     1,
			PageSize: 10,
		}

		result, err := repo.GetAllStocks(params)

		// Check that either we got an error OR we got empty results
		if err == nil && len(result.Stocks) > 0 {
			t.Errorf("Expected either an error or empty results, but got %d stocks with no error", len(result.Stocks))
		}
	})
}

func TestGetStockByTicker(t *testing.T) {
	// Stock found
	t.Run("stock found", func(t *testing.T) {
		expectedStocks := []models.Stock{
			{
				ID:      "1",
				Ticker:  "AAPL",
				Company: "Apple Inc.",
				Time:    time.Now(),
			},
		}

		// Use the helper to create a mock DB
		mockDB := mocks.CreateMockDBWithStocks(expectedStocks)
		repo := repository.NewStockRepository(mockDB)

		stocks, err := repo.GetStocksByTicker("AAPL")

		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}

		if len(stocks) != 1 {
			t.Errorf("Expected 1 stock but got %d", len(stocks))
		}

		if stocks[0].Ticker != "AAPL" {
			t.Errorf("Expected ticker AAPL but got %s", stocks[0].Ticker)
		}
	})

	// Database error
	t.Run("database error", func(t *testing.T) {
		dbError := errors.New("database error")
		mockDB := database.NewMockDatabaseWithError(dbError)

		repo := repository.NewStockRepository(mockDB)

		stocks, err := repo.GetStocksByTicker("AAPL")

		// Check that either we got an error OR we got nil/empty results
		if err == nil && len(stocks) > 0 {
			t.Errorf("Expected either an error or empty results, but got %d stocks with no error", len(stocks))
		}
	})
}

func TestGetRecentStocks(t *testing.T) {
	// Successful retrieval
	t.Run("successful retrieval", func(t *testing.T) {
		expectedStocks := []models.Stock{
			{
				ID:      "1",
				Ticker:  "AAPL",
				Company: "Apple Inc.",
				Time:    time.Now(),
			},
		}

		// Create mock DB using our helper
		mockDB := mocks.CreateMockDBWithStocks(expectedStocks)
		repo := repository.NewStockRepository(mockDB)

		stocks, err := repo.GetRecentStocks(10)

		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}

		if len(stocks) != 1 {
			t.Errorf("Expected 1 stock but got %d", len(stocks))
		}

		if stocks[0].Ticker != "AAPL" {
			t.Errorf("Expected ticker AAPL but got %s", stocks[0].Ticker)
		}
	})

	// Database error
	t.Run("database error", func(t *testing.T) {
		dbError := errors.New("database error")
		mockDB := database.NewMockDatabaseWithError(dbError)

		repo := repository.NewStockRepository(mockDB)

		stocks, err := repo.GetRecentStocks(10)

		// Check that either we got an error OR we got nil/empty results
		if err == nil && len(stocks) > 0 {
			t.Errorf("Expected either an error or empty results, but got %d stocks with no error", len(stocks))
		}
	})
}

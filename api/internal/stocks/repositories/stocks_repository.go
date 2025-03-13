package repository

import (
	"fmt"
	"stonks-api/cmd/database"
	"stonks-api/internal/stocks/models"
	"time"
)

type StockRepository struct {
	db database.Database
}

func NewStockRepository(db database.Database) *StockRepository {
	return &StockRepository{
		db: db,
	}
}

// SaveStocks saves a batch of stock data to the database
func (r *StockRepository) SaveStocks(stocks []models.Stock) error {
	if len(stocks) == 0 {
		return nil
	}

	// Process the entire batch in a single transaction
	err := r.db.Transaction(func(tx database.Transaction) error {
		for _, stock := range stocks {
			// Try to find an existing record using index-optimized query
			var count int64
			query := tx.Model(&models.Stock{}).Where("ticker = ? AND time = ?", stock.Ticker, stock.Time)

			count, err := query.Count()
			if err != nil {
				return err
			}

			if count == 0 {
				// Record doesn't exist, create it
				if err := tx.Create(&stock); err != nil {
					return err
				}
			} else {
				// Record exists, update only the necessary fields
				updates := map[string]interface{}{
					"company":     stock.Company,
					"brokerage":   stock.Brokerage,
					"action":      stock.Action,
					"rating_from": stock.RatingFrom,
					"rating_to":   stock.RatingTo,
					"target_from": stock.TargetFrom,
					"target_to":   stock.TargetTo,
					"updated_at":  time.Now(),
				}

				if err := tx.Model(&models.Stock{}).Where("ticker = ? AND time = ?",
					stock.Ticker, stock.Time).Updates(updates); err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to save stock batch: %w", err)
	}

	return nil
}

// GetAllStocks retrieves all stocks from the database with pagination
func (r *StockRepository) GetAllStocks(params models.PaginationParams) (models.PaginatedStocks, error) {
	page := params.Page
	pageSize := params.PageSize

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	totalCount, err := r.db.Count(&models.Stock{})
	if err != nil {
		return models.PaginatedStocks{}, fmt.Errorf("failed to get stock count: %w", err)
	}

	totalPages := int((totalCount + int64(pageSize) - 1) / int64(pageSize))

	var stocks []models.Stock
	err = r.db.Select("id, ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time, updated_at").
		Order("time DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&stocks)

	if err != nil {
		return models.PaginatedStocks{}, fmt.Errorf("failed to retrieve stocks: %w", err)
	}

	return models.PaginatedStocks{
		Stocks:     stocks,
		TotalCount: totalCount,
		PageSize:   pageSize,
		Page:       page,
		TotalPages: totalPages,
	}, nil
}

// GetStocksByTicker retrieves stocks by ticker with optimized query
func (r *StockRepository) GetStocksByTicker(ticker string) ([]models.Stock, error) {
	var stocks []models.Stock

	err := r.db.Select("id, ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time, updated_at").
		Where("ticker = ?", ticker).
		Order("time DESC").
		Find(&stocks)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve stocks for ticker %s: %w", ticker, err)
	}

	return stocks, nil
}

// GetRecentStocks retrieves the most recent stocks up to the limit
func (r *StockRepository) GetRecentStocks(limit int) ([]models.Stock, error) {
	var stocks []models.Stock

	// Select only the fields needed for recommendations
	err := r.db.Select("id, ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time").
		Order("time DESC").
		Limit(limit).
		Find(&stocks)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve recent stocks: %w", err)
	}

	return stocks, nil
}

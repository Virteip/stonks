package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"stonks-api/internal/stocks/models"
	"strconv"
	"strings"
	"time"
)

// StockRepository defines the interface for stock data storage
type StockRepository interface {
	SaveStocks(stocks []models.Stock) error
	GetAllStocks(params models.PaginationParams) (models.PaginatedStocks, error)
	GetStocksByTicker(ticker string) ([]models.Stock, error)
	GetRecentStocks(limit int) ([]models.Stock, error)
}

// APIConfig holds the configuration for the external API
type ExternalAPIConfig struct {
	URL        string
	AuthHeader string
	AuthToken  string
}

// StockResponse represents the API response format
type StockResponse struct {
	Items    []StockItem `json:"items"`
	NextPage string      `json:"next_page"`
}

// StockItem represents each item in the API response
type StockItem struct {
	Ticker     string    `json:"ticker"`
	Company    string    `json:"company"`
	Brokerage  string    `json:"brokerage"`
	Action     string    `json:"action"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	TargetFrom string    `json:"target_from"`
	TargetTo   string    `json:"target_to"`
	Time       time.Time `json:"time"`
}

// HTTPClient defines the interface for making HTTP requests
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type StockService struct {
	httpClient        HTTPClient
	repository        StockRepository
	externalAPIConfig ExternalAPIConfig
}

// NewStockService creates a new instance of StockService
func NewStockService(repository StockRepository) *StockService {
	return &StockService{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		repository: repository,
	}
}

// SetHTTPClient allows setting a custom HTTP client (useful for testing)
func (s *StockService) SetHTTPClient(client HTTPClient) {
	s.httpClient = client
}

// FetchStocks retrieves stock data from the API
func (s *StockService) FetchStocks(nextPage string) (*StockResponse, error) {
	url := s.externalAPIConfig.URL
	if nextPage != "" {
		url = fmt.Sprintf("%s?next_page=%s", url, nextPage)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add(s.externalAPIConfig.AuthHeader, s.externalAPIConfig.AuthToken)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("external stocks api request failed: %d, body: %s", resp.StatusCode, string(body))
	}

	var response StockResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}

// ConvertToStocks converts API items to Stock models
func (s *StockService) ConvertToStocks(items []StockItem) []models.Stock {
	stocks := make([]models.Stock, 0, len(items))

	for _, item := range items {
		// Parse the item to get proper float values
		parsedItem := s.parseStockItem(item)

		stock := models.Stock{
			Ticker:     parsedItem.Ticker,
			Company:    parsedItem.Company,
			Brokerage:  parsedItem.Brokerage,
			Action:     parsedItem.Action,
			RatingFrom: parsedItem.RatingFrom,
			RatingTo:   parsedItem.RatingTo,
			TargetFrom: parsedItem.TargetFrom,
			TargetTo:   parsedItem.TargetTo,
			Time:       parsedItem.Time,
			UpdatedAt:  time.Now(),
		}

		stocks = append(stocks, stock)
	}

	return stocks
}

// parseStockItem converts a StockItem to Stock with proper type conversions
type ParsedStockItem struct {
	Ticker     string    `json:"ticker"`
	Company    string    `json:"company"`
	Brokerage  string    `json:"brokerage"`
	Action     string    `json:"action"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	TargetFrom float64   `json:"target_from"`
	TargetTo   float64   `json:"target_to"`
	Time       time.Time `json:"time"`
}

// parseStockItem converts string target values to float64
func (s *StockService) parseStockItem(item StockItem) ParsedStockItem {
	return ParsedStockItem{
		Ticker:     item.Ticker,
		Company:    item.Company,
		Brokerage:  item.Brokerage,
		Action:     item.Action,
		RatingFrom: item.RatingFrom,
		RatingTo:   item.RatingTo,
		TargetFrom: parseTargetValue(item.TargetFrom),
		TargetTo:   parseTargetValue(item.TargetTo),
		Time:       item.Time,
	}
}

// parseTargetValue converts a price string like "$33.00" to a float64 value 33.0
func parseTargetValue(val string) float64 {
	// Remove the dollar sign
	val = strings.TrimPrefix(val, "$")

	// Parse the string to a float
	value, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0
	}

	return value
}

// SyncStocks fetches stocks from API and saves them in batches
func (s *StockService) SyncStocks() (int, error) {
	if s.externalAPIConfig.URL == "" {
		return 0, fmt.Errorf("external API URL not configured")
	}

	totalCount := 0
	batchSize := 100
	batch := make([]models.Stock, 0, batchSize)
	nextPage := ""

	fmt.Println("Starting to sync stocks from external API")

	for {
		response, err := s.FetchStocks(nextPage)
		if err != nil {
			return totalCount, fmt.Errorf("error fetching stocks: %w", err)
		}

		stocks := s.ConvertToStocks(response.Items)

		batch = append(batch, stocks...)

		for len(batch) >= batchSize {
			fmt.Printf("Saving batch of %d stocks (total processed: %d)\n", batchSize, totalCount)
			if err := s.repository.SaveStocks(batch[:batchSize]); err != nil {
				return totalCount, fmt.Errorf("error saving stocks batch: %w", err)
			}

			totalCount += batchSize
			batch = batch[batchSize:]
		}

		if response.NextPage == "" {
			break
		}

		nextPage = response.NextPage
	}

	// Save any remaining stocks
	if len(batch) > 0 {
		fmt.Printf("Saving final batch of %d stocks\n", len(batch))
		if err := s.repository.SaveStocks(batch); err != nil {
			return totalCount, fmt.Errorf("error saving final batch: %w", err)
		}
		totalCount += len(batch)
	}

	fmt.Printf("Successfully synced %d stocks from external API\n", totalCount)
	return totalCount, nil
}

// GetAllStocks retrieves all stocks with pagination
func (s *StockService) GetAllStocks(page, pageSize int) (models.PaginatedStocks, error) {
	params := models.PaginationParams{
		Page:     page,
		PageSize: pageSize,
	}
	return s.repository.GetAllStocks(params)
}

// GetStocksByTicker retrieves stocks for a specific ticker
func (s *StockService) GetStocksByTicker(ticker string) ([]models.Stock, error) {
	return s.repository.GetStocksByTicker(ticker)
}

// SetExternalAPIConfig sets the external API configuration
func (s *StockService) SetExternalAPIConfig(config ExternalAPIConfig) {
	s.externalAPIConfig = config
}

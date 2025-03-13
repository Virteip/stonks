package handlers

import (
	"net/http"
	"stonks-api/internal/stocks/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type StockHandler struct {
	stockService *services.StockService
}

func NewStockHandler(stockService *services.StockService) *StockHandler {
	return &StockHandler{
		stockService: stockService,
	}
}

// SyncStocks handles the API endpoint to fetch and store stock data
func (h *StockHandler) SyncStocks(c echo.Context) error {
	count, err := h.stockService.SyncStocks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to sync stocks: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully synced stocks",
		"count":   count,
	})
}

// GetAllStocks handles the API endpoint to retrieve all stocks with pagination
func (h *StockHandler) GetAllStocks(c echo.Context) error {
	// Parse pagination parameters
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	paginatedStocks, err := h.stockService.GetAllStocks(page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve stocks: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, paginatedStocks)
}

// GetStockByTicker handles the API endpoint to retrieve a stock by ticker
func (h *StockHandler) GetStockByTicker(c echo.Context) error {
	ticker := c.Param("ticker")
	if ticker == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Ticker parameter is required",
		})
	}

	stocks, err := h.stockService.GetStocksByTicker(ticker)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve stock: " + err.Error(),
		})
	}

	if len(stocks) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "No stock found with ticker: " + ticker,
		})
	}

	return c.JSON(http.StatusOK, stocks)
}

// RegisterRoutes registers the stock routes with the Echo router
func (h *StockHandler) RegisterRoutes(e *echo.Group) {
	e.GET("/stocks", h.GetAllStocks)
	e.GET("/stock/:ticker", h.GetStockByTicker)
	e.POST("/refresh-stocks", h.SyncStocks)
}

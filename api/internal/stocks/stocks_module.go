package stocks

import (
	"stonks-api/cmd/database"
	"stonks-api/internal/stocks/handlers"
	repository "stonks-api/internal/stocks/repositories"
	"stonks-api/internal/stocks/services"

	"github.com/labstack/echo/v4"
)

type Module struct {
	StockHandler *handlers.StockHandler
	StockService *services.StockService
}

func NewModule(db database.Database) *Module {
	// Create repository
	stockRepo := repository.NewStockRepository(db)

	// Create services
	stockService := services.NewStockService(stockRepo)

	// Create handlers
	stockHandler := handlers.NewStockHandler(stockService)

	return &Module{
		StockHandler: stockHandler,
		StockService: stockService,
	}
}

func (m *Module) RegisterRoutes(e *echo.Group) {
	m.StockHandler.RegisterRoutes(e)
}

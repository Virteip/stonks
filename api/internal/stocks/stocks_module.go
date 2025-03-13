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
	stockRepo := repository.NewStockRepository(db)
	stockService := services.NewStockService(stockRepo)
	stockHandler := handlers.NewStockHandler(stockService)

	return &Module{
		StockHandler: stockHandler,
		StockService: stockService,
	}
}

func (m *Module) RegisterRoutes(e *echo.Group) {
	m.StockHandler.RegisterRoutes(e)
}

package recommendations

import (
	"stonks-api/cmd/database"
	"stonks-api/internal/recommendations/handlers"
	"stonks-api/internal/recommendations/services"
	stocksRepository "stonks-api/internal/stocks/repositories"

	"github.com/labstack/echo/v4"
)

type Module struct {
	RecommendationHandler *handlers.RecommendationHandler
	RecommendationService *services.RecommendationService
}

func NewModule(db database.Database) *Module {
	// The recommendations module depends on the stocks repository
	stockRepo := stocksRepository.NewStockRepository(db)

	// Create the recommendation service
	recommendationService := services.NewRecommendationService(stockRepo)

	// Create the recommendation handler
	recommendationHandler := handlers.NewRecommendationHandler(recommendationService)

	return &Module{
		RecommendationHandler: recommendationHandler,
		RecommendationService: recommendationService,
	}
}

func (m *Module) RegisterRoutes(e *echo.Group) {
	m.RecommendationHandler.RegisterRoutes(e)
}

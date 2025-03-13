package handlers

import (
	"net/http"
	"stonks-api/internal/recommendations/services"

	"github.com/labstack/echo/v4"
)

type RecommendationHandler struct {
	recommendationService services.RecommendationServiceInterface
}

func NewRecommendationHandler(recommendationService services.RecommendationServiceInterface) *RecommendationHandler {
	return &RecommendationHandler{
		recommendationService: recommendationService,
	}
}

func (h *RecommendationHandler) GetRecommendations(c echo.Context) error {
	recommendations, err := h.recommendationService.GetRecommendations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get recommendations: " + err.Error(),
		})
	}

	if len(recommendations) == 0 {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "No recommendations available at this time",
		})
	}

	return c.JSON(http.StatusOK, recommendations)
}

func (h *RecommendationHandler) RegisterRoutes(e *echo.Group) {
	e.GET("/recommendations", h.GetRecommendations)
}

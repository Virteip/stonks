package services_test

import (
	"errors"
	"stonks-api/internal/recommendations/mocks"
	"stonks-api/internal/recommendations/services"
	"stonks-api/internal/stocks/models"
	"testing"
	"time"
)

func TestGetRecommendations(t *testing.T) {
	// Successful recommendations
	t.Run("successful recommendations", func(t *testing.T) {
		// Create mock stocks with various scenarios
		stocks := []models.Stock{
			{
				Ticker:     "AAPL",
				Company:    "Apple Inc.",
				Brokerage:  "Example Brokerage",
				Action:     "upgraded by",
				RatingFrom: "Hold",
				RatingTo:   "Buy",
				TargetFrom: 150.0,
				TargetTo:   200.0,
				Time:       time.Now(),
			},
			{
				Ticker:     "MSFT",
				Company:    "Microsoft Corp",
				Brokerage:  "Example Brokerage",
				Action:     "target raised by",
				RatingFrom: "Buy",
				RatingTo:   "Buy",
				TargetFrom: 300.0,
				TargetTo:   320.0,
				Time:       time.Now(),
			},
			{
				Ticker:     "GOOG",
				Company:    "Alphabet Inc.",
				Brokerage:  "Example Brokerage",
				Action:     "downgraded by",
				RatingFrom: "Buy",
				RatingTo:   "Sell",
				TargetFrom: 120.0,
				TargetTo:   90.0,
				Time:       time.Now(),
			},
		}

		mockRepo := &mocks.MockStockRepository{
			GetRecentStocksFn: func(limit int) ([]models.Stock, error) {
				return stocks, nil
			},
		}

		service := services.NewRecommendationService(mockRepo)

		recommendations, err := service.GetRecommendations()

		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}

		// We should have at least one recommendation (likely AAPL and MSFT, but not GOOG as it was downgraded)
		if len(recommendations) < 1 {
			t.Errorf("Expected at least 1 recommendation but got %d", len(recommendations))
		}

		// Verify the recommendations are sorted properly (highest score first)
		if len(recommendations) > 1 {
			for i := 0; i < len(recommendations)-1; i++ {
				if recommendations[i].Score < recommendations[i+1].Score {
					t.Errorf("Recommendations not sorted correctly at position %d", i)
				}
			}
		}

		// Verify AAPL (upgraded) has a positive score
		found := false
		for _, rec := range recommendations {
			if rec.Stock.Ticker == "AAPL" {
				found = true
				if rec.Score <= 0 {
					t.Errorf("Expected positive score for AAPL but got %f", rec.Score)
				}
				if rec.Reason == "" {
					t.Errorf("Expected recommendation reason to be non-empty")
				}
				break
			}
		}

		if !found && len(recommendations) > 0 {
			t.Errorf("Expected to find AAPL in recommendations")
		}
	})

	// No recommendations
	t.Run("no recommendations", func(t *testing.T) {
		// Create mock stocks with only negative scenarios
		stocks := []models.Stock{
			{
				Ticker:     "BAD1",
				Company:    "Bad Stock 1",
				Brokerage:  "Example Brokerage",
				Action:     "downgraded by",
				RatingFrom: "Neutral",
				RatingTo:   "Sell",
				TargetFrom: 50.0,
				TargetTo:   30.0,
				Time:       time.Now(),
			},
			{
				Ticker:     "BAD2",
				Company:    "Bad Stock 2",
				Brokerage:  "Example Brokerage",
				Action:     "target lowered by",
				RatingFrom: "Underperform",
				RatingTo:   "Underperform",
				TargetFrom: 40.0,
				TargetTo:   20.0,
				Time:       time.Now(),
			},
		}

		mockRepo := &mocks.MockStockRepository{
			GetRecentStocksFn: func(limit int) ([]models.Stock, error) {
				return stocks, nil
			},
		}

		service := services.NewRecommendationService(mockRepo)

		recommendations, err := service.GetRecommendations()

		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}

		if len(recommendations) != 0 {
			t.Errorf("Expected 0 recommendations but got %d", len(recommendations))
		}
	})

	// Repository error
	t.Run("repository error", func(t *testing.T) {
		mockRepo := &mocks.MockStockRepository{
			GetRecentStocksFn: func(limit int) ([]models.Stock, error) {
				return nil, errors.New("repository error")
			},
		}

		service := services.NewRecommendationService(mockRepo)

		_, err := service.GetRecommendations()

		if err == nil {
			t.Errorf("Expected error but got nil")
		}
	})

	// Test with multiple positive stock ratings
	t.Run("multiple positive stocks", func(t *testing.T) {
		// Create mock stocks with various positive scenarios
		stocks := []models.Stock{
			{
				Ticker:     "AAPL",
				Company:    "Apple Inc.",
				Brokerage:  "Brokerage 1",
				Action:     "upgraded by",
				RatingFrom: "Hold",
				RatingTo:   "Buy",
				TargetFrom: 150.0,
				TargetTo:   200.0,
				Time:       time.Now(),
			},
			{
				Ticker:     "MSFT",
				Company:    "Microsoft Corp",
				Brokerage:  "Brokerage 1",
				Action:     "target raised by",
				RatingFrom: "Buy",
				RatingTo:   "Buy",
				TargetFrom: 300.0,
				TargetTo:   350.0,
				Time:       time.Now(),
			},
			{
				Ticker:     "GOOGL",
				Company:    "Alphabet Inc.",
				Brokerage:  "Brokerage 2",
				Action:     "reiterated by",
				RatingFrom: "Outperform",
				RatingTo:   "Outperform",
				TargetFrom: 150.0,
				TargetTo:   155.0,
				Time:       time.Now(),
			},
			{
				Ticker:     "AMZN",
				Company:    "Amazon.com Inc.",
				Brokerage:  "Brokerage 3",
				Action:     "upgraded by",
				RatingFrom: "Neutral",
				RatingTo:   "Overweight",
				TargetFrom: 170.0,
				TargetTo:   185.0,
				Time:       time.Now(),
			},
			{
				Ticker:     "FB",
				Company:    "Meta Platforms Inc.",
				Brokerage:  "Brokerage 1",
				Action:     "upgraded by",
				RatingFrom: "Neutral",
				RatingTo:   "Buy",
				TargetFrom: 320.0,
				TargetTo:   380.0,
				Time:       time.Now(),
			},
			{
				Ticker:     "NFLX",
				Company:    "Netflix Inc.",
				Brokerage:  "Brokerage 2",
				Action:     "target raised by",
				RatingFrom: "Outperform",
				RatingTo:   "Outperform",
				TargetFrom: 600.0,
				TargetTo:   650.0,
				Time:       time.Now(),
			},
		}

		mockRepo := &mocks.MockStockRepository{
			GetRecentStocksFn: func(limit int) ([]models.Stock, error) {
				return stocks, nil
			},
		}

		service := services.NewRecommendationService(mockRepo)

		recommendations, err := service.GetRecommendations()

		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}

		// Should return exactly 5 recommendations as per the service implementation
		if len(recommendations) != 5 {
			t.Errorf("Expected 5 recommendations but got %d", len(recommendations))
		}

		// Verify the recommendations are sorted properly (highest score first)
		for i := 0; i < len(recommendations)-1; i++ {
			if recommendations[i].Score < recommendations[i+1].Score {
				t.Errorf("Recommendations not sorted correctly at position %d", i)
			}
		}
	})
}

func TestGetRatingScore(t *testing.T) {
	t.Run("test rating score calculation", func(t *testing.T) {
		// Test different rating values to make sure they map to the right categories
		positiveRatings := []string{"Buy", "Strong-Buy", "Outperform", "Overweight"}
		neutralRatings := []string{"Hold", "Neutral", "Equal Weight", "Market Perform"}
		negativeRatings := []string{"Sell", "Reduce", "Underperform", "Underweight"}

		for _, rating := range positiveRatings {
			score := services.GetRatingScore(rating)
			if score != 5 {
				t.Errorf("Expected positive score of 5 for rating %s but got %d", rating, score)
			}
		}

		for _, rating := range neutralRatings {
			score := services.GetRatingScore(rating)
			if score != 3 {
				t.Errorf("Expected neutral score of 3 for rating %s but got %d", rating, score)
			}
		}

		for _, rating := range negativeRatings {
			score := services.GetRatingScore(rating)
			if score != 1 {
				t.Errorf("Expected negative score of 1 for rating %s but got %d", rating, score)
			}
		}

		// Test unknown rating
		score := services.GetRatingScore("Unknown Rating")
		if score != 3 {
			t.Errorf("Expected neutral score of 3 for unknown rating but got %d", score)
		}
	})
}

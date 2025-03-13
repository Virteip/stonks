package services

import (
	"sort"
	"stonks-api/internal/stocks/models"
)

type StockRepository interface {
	GetRecentStocks(limit int) ([]models.Stock, error)
}

type RecommendationServiceInterface interface {
	GetRecommendations() ([]StockRecommendation, error)
}

type StockRecommendation struct {
	Stock  models.Stock `json:"stock"`
	Score  float64      `json:"score"`
	Reason string       `json:"reason"`
}

type RecommendationService struct {
	stockRepository StockRepository
}

func NewRecommendationService(stockRepository StockRepository) *RecommendationService {
	return &RecommendationService{
		stockRepository: stockRepository,
	}
}

func (s *RecommendationService) GetRecommendations() ([]StockRecommendation, error) {
	// Only fetch the 200 most recent stocks instead of all stocks
	stocks, err := s.stockRepository.GetRecentStocks(200)
	if err != nil {
		return nil, err
	}

	recommendations := make([]StockRecommendation, 0, len(stocks)/2)
	// Map to ensure to only include one recommendation per ticker
	tickerMap := make(map[string]bool)

	for _, stock := range stocks {
		if _, exists := tickerMap[stock.Ticker]; exists {
			continue
		}

		score, reason := s.calculateScore(stock)

		if score > 0 {
			recommendations = append(recommendations, StockRecommendation{
				Stock:  stock,
				Score:  score,
				Reason: reason,
			})

			// Mark ticker as processed
			tickerMap[stock.Ticker] = true
		}
	}

	// Sort by score (highest first)
	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].Score > recommendations[j].Score
	})

	// Return top 5 recommendations or all if less than 5
	if len(recommendations) > 5 {
		return recommendations[:5], nil
	}
	return recommendations, nil
}

// calculateScore assigns a score to a stock based on various factors
func (s *RecommendationService) calculateScore(stock models.Stock) (float64, string) {
	var score float64
	var reason string

	// 1: Upgrade vs downgrade
	if stock.Action == "upgraded by" {
		score += 2.0
		reason = "Stock was recently upgraded"
	} else if stock.Action == "downgraded by" {
		score -= 2.0
		reason = "Stock was recently downgraded"
	}

	// 2: Target price change
	targetChange := stock.TargetTo - stock.TargetFrom
	targetPercentChange := 0.0
	if stock.TargetFrom > 0 {
		targetPercentChange = (targetChange / stock.TargetFrom) * 100
	}

	if targetPercentChange > 10 {
		score += 2.0
		if reason == "" {
			reason = "Target price increased significantly"
		} else {
			reason += ", Target price increased significantly"
		}
	} else if targetPercentChange > 0 {
		score += 1.0
		if reason == "" {
			reason = "Target price increased"
		} else {
			reason += ", Target price increased"
		}
	} else if targetPercentChange < -10 {
		score -= 2.0
		if reason == "" {
			reason = "Target price decreased significantly"
		} else {
			reason += ", Target price decreased significantly"
		}
	} else if targetPercentChange < 0 {
		score -= 1.0
		if reason == "" {
			reason = "Target price decreased"
		} else {
			reason += ", Target price decreased"
		}
	}

	// 3: Rating improvement
	fromScore := GetRatingScore(stock.RatingFrom)
	toScore := GetRatingScore(stock.RatingTo)

	ratingChange := toScore - fromScore
	if ratingChange > 0 {
		score += float64(ratingChange) * 0.5
		if reason == "" {
			reason = "Rating improved"
		} else {
			reason += ", Rating improved"
		}
	} else if ratingChange < 0 {
		score += float64(ratingChange) * 0.5
		if reason == "" {
			reason = "Rating downgraded"
		} else {
			reason += ", Rating downgraded"
		}
	} else if toScore >= 5 {
		score += 0.5
		if reason == "" {
			reason = "Maintained positive rating"
		} else {
			reason += ", Maintained positive rating"
		}
	}

	// 4: Current rating strength
	if toScore >= 7 { // Strong Buy or Buy
		score += 2.0
		if reason == "" {
			reason = "Strong positive rating"
		} else {
			reason += ", Strong positive rating"
		}
	} else if toScore >= 5 { // Outperform, Overweight
		score += 1.0
		if reason == "" {
			reason = "Positive rating"
		} else {
			reason += ", Positive rating"
		}
	}

	return score, reason
}

package services

// Rating categories for recommendation calculation
const (
	RatingCategoryPositive = "Positive"
	RatingCategoryNeutral  = "Neutral"
	RatingCategoryNegative = "Negative"
)

// getRatingCategoryMap maps specific rating strings to their categories
func getRatingCategoryMap() map[string]string {
	return map[string]string{
		// Positive ratings
		"Buy":               RatingCategoryPositive,
		"Strong-Buy":        RatingCategoryPositive,
		"Outperform":        RatingCategoryPositive,
		"Outperformer":      RatingCategoryPositive,
		"Overweight":        RatingCategoryPositive,
		"Positive":          RatingCategoryPositive,
		"Market Outperform": RatingCategoryPositive,
		"Sector Outperform": RatingCategoryPositive,

		// Neutral ratings
		"Hold":           RatingCategoryNeutral,
		"Neutral":        RatingCategoryNeutral,
		"Equal Weight":   RatingCategoryNeutral,
		"Market Perform": RatingCategoryNeutral,
		"Sector Perform": RatingCategoryNeutral,
		"In-Line":        RatingCategoryNeutral,
		"Inline":         RatingCategoryNeutral,
		"Peer Perform":   RatingCategoryNeutral,
		"Sector Weight":  RatingCategoryNeutral,

		// Negative ratings
		"Sell":                RatingCategoryNegative,
		"Reduce":              RatingCategoryNegative,
		"Underperform":        RatingCategoryNegative,
		"Underweight":         RatingCategoryNegative,
		"Negative":            RatingCategoryNegative,
		"Sector Underperform": RatingCategoryNegative,
	}
}

// ratingCategories contains the mapping from rating string to category
var ratingCategories = getRatingCategoryMap()

// GetRatingScore returns a numeric score for a rating
func GetRatingScore(rating string) int {
	category := GetRatingCategory(rating)

	switch category {
	case RatingCategoryPositive:
		return 5
	case RatingCategoryNegative:
		return 1
	default:
		return 3
	}
}

// GetRatingCategory returns the category of a rating
func GetRatingCategory(rating string) string {
	category, exists := ratingCategories[rating]
	if !exists {
		return RatingCategoryNeutral
	}
	return category
}

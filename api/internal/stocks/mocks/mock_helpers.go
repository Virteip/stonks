package mocks

import (
	"stonks-api/cmd/database"
	"stonks-api/internal/stocks/models"
)

// CreateMockQueryChain creates a mock query chain that returns the specified stocks
func CreateMockQueryChain(stocks []models.Stock) database.Query {
	return &database.MockQuery{
		FindFn: func(dest interface{}, conditions ...interface{}) error {
			// Populate the destination with our expected stocks
			stocksPtr := dest.(*[]models.Stock)
			*stocksPtr = stocks
			return nil
		},
		OrderFn: func(value interface{}) database.Query {
			return &database.MockQuery{
				FindFn: func(dest interface{}, conditions ...interface{}) error {
					stocksPtr := dest.(*[]models.Stock)
					*stocksPtr = stocks
					return nil
				},
				LimitFn: func(limit int) database.Query {
					return &database.MockQuery{
						FindFn: func(dest interface{}, conditions ...interface{}) error {
							stocksPtr := dest.(*[]models.Stock)
							*stocksPtr = stocks
							return nil
						},
						OffsetFn: func(offset int) database.Query {
							return &database.MockQuery{
								FindFn: func(dest interface{}, conditions ...interface{}) error {
									stocksPtr := dest.(*[]models.Stock)
									*stocksPtr = stocks
									return nil
								},
							}
						},
					}
				},
			}
		},
		WhereFn: func(query interface{}, args ...interface{}) database.Query {
			return &database.MockQuery{
				OrderFn: func(value interface{}) database.Query {
					return &database.MockQuery{
						FindFn: func(dest interface{}, conditions ...interface{}) error {
							stocksPtr := dest.(*[]models.Stock)
							*stocksPtr = stocks
							return nil
						},
					}
				},
			}
		},
	}
}

// CreateMockDBWithStocks creates a mock database that returns the specified stocks
func CreateMockDBWithStocks(stocks []models.Stock) *database.MockDatabase {
	mockQuery := CreateMockQueryChain(stocks)

	return &database.MockDatabase{
		SelectFn: func(query interface{}, args ...interface{}) database.Query {
			return mockQuery
		},
		ModelFn: func(value interface{}) database.Query {
			return mockQuery
		},
		WhereFn: func(query interface{}, args ...interface{}) database.Query {
			return mockQuery
		},
		CountFn: func(model interface{}) (int64, error) {
			return int64(len(stocks)), nil
		},
	}
}

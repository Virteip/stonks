package mocks

import (
	"net/http"
	"stonks-api/internal/stocks/models"
)

// MockRepository is a simplified mock implementation of the Repository interface
type MockRepository struct {
	StockByTicker *models.Stock
	Stocks        []models.Stock
	PaginatedData models.PaginatedStocks
	ErrorToReturn error
}

func (m *MockRepository) SaveStocks(stocks []models.Stock) error {
	return m.ErrorToReturn
}

func (m *MockRepository) GetAllStocks(page, pageSize int) (models.PaginatedStocks, error) {
	return m.PaginatedData, m.ErrorToReturn
}

func (m *MockRepository) GetStockByTicker(ticker string) (*models.Stock, error) {
	return m.StockByTicker, m.ErrorToReturn
}

func (m *MockRepository) GetRecentStocks(limit int) ([]models.Stock, error) {
	return m.Stocks, m.ErrorToReturn
}

type MockHTTPClient struct {
	Response *http.Response
	Error    error
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.Response, m.Error
}

// MockDB is a simplified mock implementation of the database operations
type MockDB struct {
	Error error
}

func (m *MockDB) Model(value interface{}) interface{}                       { return m }
func (m *MockDB) Where(query interface{}, args ...interface{}) interface{}  { return m }
func (m *MockDB) Order(value interface{}) interface{}                       { return m }
func (m *MockDB) Limit(limit int) interface{}                               { return m }
func (m *MockDB) Offset(offset int) interface{}                             { return m }
func (m *MockDB) Select(query interface{}, args ...interface{}) interface{} { return m }
func (m *MockDB) Count(count *int64) error                                  { return m.Error }
func (m *MockDB) Find(dest interface{}, conds ...interface{}) error         { return m.Error }
func (m *MockDB) First(dest interface{}, conds ...interface{}) error        { return m.Error }
func (m *MockDB) Create(value interface{}) error                            { return m.Error }
func (m *MockDB) Updates(values interface{}) error                          { return m.Error }
func (m *MockDB) Transaction(fc func(tx interface{}) error) error           { return fc(m) }

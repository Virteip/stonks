package database

// MockDatabase provides a mock implementation of the Database interface for testing
type MockDatabase struct {
	TransactionFn func(fc func(tx Transaction) error) error
	FindFn        func(dest interface{}, conditions ...interface{}) error
	CreateFn      func(value interface{}) error
	UpdatesFn     func(model interface{}, values interface{}) error
	DeleteFn      func(value interface{}) error
	WhereFn       func(query interface{}, args ...interface{}) Query
	OrderFn       func(value interface{}) Query
	LimitFn       func(limit int) Query
	OffsetFn      func(offset int) Query
	SelectFn      func(query interface{}, args ...interface{}) Query
	CountFn       func(model interface{}) (int64, error)
	CloseFn       func() error
	ModelFn       func(value interface{}) Query
	ExecFn        func(sql string, values ...interface{}) error
	PingFn        func() error
}

// Transaction runs operations within a database transaction
func (m *MockDatabase) Transaction(fc func(tx Transaction) error) error {
	if m.TransactionFn != nil {
		return m.TransactionFn(fc)
	}
	return fc(&MockTransaction{})
}

// Find retrieves records matching the given conditions
func (m *MockDatabase) Find(dest interface{}, conditions ...interface{}) error {
	if m.FindFn != nil {
		return m.FindFn(dest, conditions...)
	}
	return nil
}

// Create inserts a new record into the database
func (m *MockDatabase) Create(value interface{}) error {
	if m.CreateFn != nil {
		return m.CreateFn(value)
	}
	return nil
}

// Updates updates a record with the given values
func (m *MockDatabase) Updates(model interface{}, values interface{}) error {
	if m.UpdatesFn != nil {
		return m.UpdatesFn(model, values)
	}
	return nil
}

// Delete removes a record from the database
func (m *MockDatabase) Delete(value interface{}) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(value)
	}
	return nil
}

// Where adds a condition to the query
func (m *MockDatabase) Where(query interface{}, args ...interface{}) Query {
	if m.WhereFn != nil {
		return m.WhereFn(query, args...)
	}
	return &MockQuery{}
}

// Order specifies the order of results
func (m *MockDatabase) Order(value interface{}) Query {
	if m.OrderFn != nil {
		return m.OrderFn(value)
	}
	return &MockQuery{}
}

// Limit specifies the maximum number of records to return
func (m *MockDatabase) Limit(limit int) Query {
	if m.LimitFn != nil {
		return m.LimitFn(limit)
	}
	return &MockQuery{}
}

// Offset specifies the number of records to skip
func (m *MockDatabase) Offset(offset int) Query {
	if m.OffsetFn != nil {
		return m.OffsetFn(offset)
	}
	return &MockQuery{}
}

// Select specifies the fields to select
func (m *MockDatabase) Select(query interface{}, args ...interface{}) Query {
	if m.SelectFn != nil {
		return m.SelectFn(query, args...)
	}
	return &MockQuery{}
}

// Count returns the total number of records matching the query
func (m *MockDatabase) Count(model interface{}) (int64, error) {
	if m.CountFn != nil {
		return m.CountFn(model)
	}
	return 0, nil
}

// Close closes the database connection
func (m *MockDatabase) Close() error {
	if m.CloseFn != nil {
		return m.CloseFn()
	}
	return nil
}

// Model specifies the model for the query
func (m *MockDatabase) Model(value interface{}) Query {
	if m.ModelFn != nil {
		return m.ModelFn(value)
	}
	return &MockQuery{}
}

// Exec executes raw SQL
func (m *MockDatabase) Exec(sql string, values ...interface{}) error {
	if m.ExecFn != nil {
		return m.ExecFn(sql, values...)
	}
	return nil
}

// Ping checks database connectivity
func (m *MockDatabase) Ping() error {
	if m.PingFn != nil {
		return m.PingFn()
	}
	return nil
}

// MockQuery provides a mock implementation of the Query interface for testing
type MockQuery struct {
	FindFn    func(dest interface{}, conditions ...interface{}) error
	WhereFn   func(query interface{}, args ...interface{}) Query
	OrderFn   func(value interface{}) Query
	LimitFn   func(limit int) Query
	OffsetFn  func(offset int) Query
	SelectFn  func(query interface{}, args ...interface{}) Query
	CountFn   func() (int64, error)
	UpdatesFn func(values interface{}) error
}

// Find retrieves records matching the query conditions
func (m *MockQuery) Find(dest interface{}, conditions ...interface{}) error {
	if m.FindFn != nil {
		return m.FindFn(dest, conditions...)
	}
	return nil
}

// Where adds a condition to the query
func (m *MockQuery) Where(query interface{}, args ...interface{}) Query {
	if m.WhereFn != nil {
		return m.WhereFn(query, args...)
	}
	return m
}

// Order specifies the order of results
func (m *MockQuery) Order(value interface{}) Query {
	if m.OrderFn != nil {
		return m.OrderFn(value)
	}
	return m
}

// Limit specifies the maximum number of records to return
func (m *MockQuery) Limit(limit int) Query {
	if m.LimitFn != nil {
		return m.LimitFn(limit)
	}
	return m
}

// Offset specifies the number of records to skip
func (m *MockQuery) Offset(offset int) Query {
	if m.OffsetFn != nil {
		return m.OffsetFn(offset)
	}
	return m
}

// Select specifies the fields to select
func (m *MockQuery) Select(query interface{}, args ...interface{}) Query {
	if m.SelectFn != nil {
		return m.SelectFn(query, args...)
	}
	return m
}

// Count returns the total number of records matching the query
func (m *MockQuery) Count() (int64, error) {
	if m.CountFn != nil {
		return m.CountFn()
	}
	return 0, nil
}

// Updates updates records with the given values
func (m *MockQuery) Updates(values interface{}) error {
	if m.UpdatesFn != nil {
		return m.UpdatesFn(values)
	}
	return nil
}

// MockTransaction provides a mock implementation of the Transaction interface for testing
type MockTransaction struct {
	CommitFn   func() error
	RollbackFn func() error
	FindFn     func(dest interface{}, conditions ...interface{}) error
	CreateFn   func(value interface{}) error
	UpdatesFn  func(model interface{}, values interface{}) error
	WhereFn    func(query interface{}, args ...interface{}) Query
	ModelFn    func(value interface{}) Query
}

// Commit commits the transaction
func (m *MockTransaction) Commit() error {
	if m.CommitFn != nil {
		return m.CommitFn()
	}
	return nil
}

// Rollback aborts the transaction
func (m *MockTransaction) Rollback() error {
	if m.RollbackFn != nil {
		return m.RollbackFn()
	}
	return nil
}

// Find retrieves records matching the given conditions
func (m *MockTransaction) Find(dest interface{}, conditions ...interface{}) error {
	if m.FindFn != nil {
		return m.FindFn(dest, conditions...)
	}
	return nil
}

// Create inserts a new record into the database
func (m *MockTransaction) Create(value interface{}) error {
	if m.CreateFn != nil {
		return m.CreateFn(value)
	}
	return nil
}

// Updates updates a record with the given values
func (m *MockTransaction) Updates(model interface{}, values interface{}) error {
	if m.UpdatesFn != nil {
		return m.UpdatesFn(model, values)
	}
	return nil
}

// Where adds a condition to the query
func (m *MockTransaction) Where(query interface{}, args ...interface{}) Query {
	if m.WhereFn != nil {
		return m.WhereFn(query, args...)
	}
	return &MockQuery{}
}

// Model specifies the model for the query
func (m *MockTransaction) Model(value interface{}) Query {
	if m.ModelFn != nil {
		return m.ModelFn(value)
	}
	return &MockQuery{}
}

// NewMockDatabaseWithError returns a mock database that returns the specified error for all operations
func NewMockDatabaseWithError(err error) *MockDatabase {
	return &MockDatabase{
		TransactionFn: func(fc func(tx Transaction) error) error {
			return err
		},
		FindFn: func(dest interface{}, conditions ...interface{}) error {
			return err
		},
		CreateFn: func(value interface{}) error {
			return err
		},
		UpdatesFn: func(model interface{}, values interface{}) error {
			return err
		},
		DeleteFn: func(value interface{}) error {
			return err
		},
		CountFn: func(model interface{}) (int64, error) {
			return 0, err
		},
		ExecFn: func(sql string, values ...interface{}) error {
			return err
		},
		PingFn: func() error {
			return err
		},
	}
}

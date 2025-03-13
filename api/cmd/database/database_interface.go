package database

// Database defines the interface for database operations
// This allows us to easily swap implementations and create mocks for testing
type Database interface {
	// Transaction runs operations within a database transaction
	Transaction(fc func(tx Transaction) error) error

	// Find retrieves records matching the given conditions
	Find(dest interface{}, conditions ...interface{}) error

	// Create inserts a new record into the database
	Create(value interface{}) error

	// Updates updates a record with the given values
	Updates(model interface{}, values interface{}) error

	// Delete removes a record from the database
	Delete(value interface{}) error

	// Where adds a condition to the query
	Where(query interface{}, args ...interface{}) Query

	// Order specifies the order of results
	Order(value interface{}) Query

	// Limit specifies the maximum number of records to return
	Limit(limit int) Query

	// Offset specifies the number of records to skip
	Offset(offset int) Query

	// Select specifies the fields to select
	Select(query interface{}, args ...interface{}) Query

	// Count returns the total number of records matching the query
	Count(model interface{}) (int64, error)

	// Close closes the database connection
	Close() error

	// Model specifies the model for the query
	Model(value interface{}) Query

	// Exec executes raw SQL
	Exec(sql string, values ...interface{}) error

	// Ping checks database connectivity
	Ping() error
}

// Query represents a database query builder
type Query interface {
	// Find retrieves records matching the query conditions
	Find(dest interface{}, conditions ...interface{}) error

	// Where adds a condition to the query
	Where(query interface{}, args ...interface{}) Query

	// Order specifies the order of results
	Order(value interface{}) Query

	// Limit specifies the maximum number of records to return
	Limit(limit int) Query

	// Offset specifies the number of records to skip
	Offset(offset int) Query

	// Select specifies the fields to select
	Select(query interface{}, args ...interface{}) Query

	// Count returns the total number of records matching the query
	Count() (int64, error)

	// Updates updates records with the given values
	Updates(values interface{}) error
}

// Transaction represents a database transaction
type Transaction interface {
	// Commit commits the transaction
	Commit() error

	// Rollback aborts the transaction
	Rollback() error

	// Find retrieves records matching the given conditions
	Find(dest interface{}, conditions ...interface{}) error

	// Create inserts a new record into the database
	Create(value interface{}) error

	// Updates updates a record with the given values
	Updates(model interface{}, values interface{}) error

	// Where adds a condition to the query
	Where(query interface{}, args ...interface{}) Query

	// Model specifies the model for the query
	Model(value interface{}) Query
}

package database

import (
	"gorm.io/gorm"
)

// GormAdapter adapts GORM's *gorm.DB to our Database interface
type GormAdapter struct {
	db *gorm.DB
}

// NewGormAdapter creates a new GORM adapter
func NewGormAdapter(db *gorm.DB) *GormAdapter {
	return &GormAdapter{db: db}
}

// Transaction runs a function within a database transaction
func (g *GormAdapter) Transaction(fc func(tx Transaction) error) error {
	return g.db.Transaction(func(tx *gorm.DB) error {
		return fc(&GormTransactionAdapter{tx: tx})
	})
}

// Find retrieves records matching the given conditions
func (g *GormAdapter) Find(dest interface{}, conditions ...interface{}) error {
	return g.db.Find(dest, conditions...).Error
}

// Create inserts a new record into the database
func (g *GormAdapter) Create(value interface{}) error {
	return g.db.Create(value).Error
}

// Updates updates a record with the given values
func (g *GormAdapter) Updates(model interface{}, values interface{}) error {
	result := g.db.Model(model).Updates(values)
	return result.Error
}

// Delete removes a record from the database
func (g *GormAdapter) Delete(value interface{}) error {
	return g.db.Delete(value).Error
}

// Where adds a condition to the query
func (g *GormAdapter) Where(query interface{}, args ...interface{}) Query {
	return &GormQueryAdapter{query: g.db.Where(query, args...)}
}

// Order specifies the order of results
func (g *GormAdapter) Order(value interface{}) Query {
	return &GormQueryAdapter{query: g.db.Order(value)}
}

// Limit specifies the maximum number of records to return
func (g *GormAdapter) Limit(limit int) Query {
	return &GormQueryAdapter{query: g.db.Limit(limit)}
}

// Offset specifies the number of records to skip
func (g *GormAdapter) Offset(offset int) Query {
	return &GormQueryAdapter{query: g.db.Offset(offset)}
}

// Select specifies the fields to select
func (g *GormAdapter) Select(query interface{}, args ...interface{}) Query {
	return &GormQueryAdapter{query: g.db.Select(query, args...)}
}

// Count returns the total number of records matching the query
func (g *GormAdapter) Count(model interface{}) (int64, error) {
	var count int64
	err := g.db.Model(model).Count(&count).Error
	return count, err
}

// Close closes the database connection
func (g *GormAdapter) Close() error {
	db, err := g.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

// Model specifies the model for the query
func (g *GormAdapter) Model(value interface{}) Query {
	return &GormQueryAdapter{query: g.db.Model(value)}
}

// Exec executes raw SQL
func (g *GormAdapter) Exec(sql string, values ...interface{}) error {
	return g.db.Exec(sql, values...).Error
}

// Ping checks database connectivity
func (g *GormAdapter) Ping() error {
	db, err := g.db.DB()
	if err != nil {
		return err
	}
	return db.Ping()
}

// GormQueryAdapter adapts GORM's query methods to our Query interface
type GormQueryAdapter struct {
	query *gorm.DB
}

// Find retrieves records matching the query conditions
func (q *GormQueryAdapter) Find(dest interface{}, conditions ...interface{}) error {
	return q.query.Find(dest, conditions...).Error
}

// Where adds a condition to the query
func (q *GormQueryAdapter) Where(query interface{}, args ...interface{}) Query {
	return &GormQueryAdapter{query: q.query.Where(query, args...)}
}

// Order specifies the order of results
func (q *GormQueryAdapter) Order(value interface{}) Query {
	return &GormQueryAdapter{query: q.query.Order(value)}
}

// Limit specifies the maximum number of records to return
func (q *GormQueryAdapter) Limit(limit int) Query {
	return &GormQueryAdapter{query: q.query.Limit(limit)}
}

// Offset specifies the number of records to skip
func (q *GormQueryAdapter) Offset(offset int) Query {
	return &GormQueryAdapter{query: q.query.Offset(offset)}
}

// Select specifies the fields to select
func (q *GormQueryAdapter) Select(query interface{}, args ...interface{}) Query {
	return &GormQueryAdapter{query: q.query.Select(query, args...)}
}

// Count returns the total number of records matching the query
func (q *GormQueryAdapter) Count() (int64, error) {
	var count int64
	err := q.query.Count(&count).Error
	return count, err
}

// GormTransactionAdapter adapts GORM's transaction methods to our Transaction interface
type GormTransactionAdapter struct {
	tx *gorm.DB
}

// Commit commits the transaction
func (t *GormTransactionAdapter) Commit() error {
	return t.tx.Commit().Error
}

// Rollback aborts the transaction
func (t *GormTransactionAdapter) Rollback() error {
	return t.tx.Rollback().Error
}

// Find retrieves records matching the given conditions
func (t *GormTransactionAdapter) Find(dest interface{}, conditions ...interface{}) error {
	return t.tx.Find(dest, conditions...).Error
}

// Create inserts a new record into the database
func (t *GormTransactionAdapter) Create(value interface{}) error {
	return t.tx.Create(value).Error
}

// Updates updates a record with the given values
func (t *GormTransactionAdapter) Updates(model interface{}, values interface{}) error {
	return t.tx.Model(model).Updates(values).Error
}

// Where adds a condition to the query
func (t *GormTransactionAdapter) Where(query interface{}, args ...interface{}) Query {
	return &GormQueryAdapter{query: t.tx.Where(query, args...)}
}

// Model specifies the model for the query
func (t *GormTransactionAdapter) Model(value interface{}) Query {
	return &GormQueryAdapter{query: t.tx.Model(value)}
}

// Updates updates records with the given values
func (q *GormQueryAdapter) Updates(values interface{}) error {
	return q.query.Updates(values).Error
}

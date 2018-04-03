package fusion

import (
	"github.com/jinzhu/gorm"
)

type QueryOpt func(*gorm.DB) (*gorm.DB, error)

// WithOrder sets the result ordering
func WithOrder(order string) QueryOpt {
	return func(db *gorm.DB) (*gorm.DB, error) {
		return db.Order(order), nil
	}
}

// WithLimit limits the number of results
func WithLimit(limit int) QueryOpt {
	return func(db *gorm.DB) (*gorm.DB, error) {
		return db.Limit(limit), nil
	}
}

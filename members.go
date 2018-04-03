package fusion

import (
	"github.com/jinzhu/gorm"
)

type ListMemberOpt func(*gorm.DB) (*gorm.DB, error)

func WithMemberOrder(order string) ListMemberOpt {
	return func(db *gorm.DB) (*gorm.DB, error) {
		return db.Order(order), nil
	}
}

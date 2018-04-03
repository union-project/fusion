package fusion

import (
	"github.com/jinzhu/gorm"
)

type ListBillOpt func(*gorm.DB) (*gorm.DB, error)

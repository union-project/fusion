package manager

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/union-project/fusion/types"
)

type Manager struct {
	Config *Config
	db     *gorm.DB
}

func NewManager(cfg *Config) (*Manager, error) {
	db, err := gorm.Open("postgres", cfg.DBAddr)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&types.Member{})

	return &Manager{
		Config: cfg,
		db:     db,
	}, nil
}

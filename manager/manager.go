package manager

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"github.com/union-project/fusion/types"
)

type Manager struct {
	Config *Config
	db     *gorm.DB
}

func NewManager(cfg *Config) (*Manager, error) {
	logrus.WithFields(logrus.Fields{
		"addr": cfg.DBAddr,
	}).Debug("initializing database")
	db, err := gorm.Open("postgres", cfg.DBAddr)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&types.Member{},
		&types.Bill{},
	)

	return &Manager{
		Config: cfg,
		db:     db,
	}, nil
}

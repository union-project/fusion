package manager

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"github.com/union-project/fusion/clients/propublica"
	"github.com/union-project/fusion/types"
)

type Manager struct {
	Config           *Config
	db               *gorm.DB
	propublicaClient *propublica.Client
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
		&types.Sponsor{},
		&types.Summary{},
	)

	client, err := propublica.NewClient(&propublica.Config{
		APIKey: cfg.ProPublicaAPIKey,
	})
	if err != nil {
		return nil, err
	}

	return &Manager{
		Config:           cfg,
		db:               db,
		propublicaClient: client,
	}, nil
}

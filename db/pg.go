package db

import (
	"time"

	"github.com/evt/video1/config"
	"github.com/go-pg/pg/v9"
	_ "github.com/lib/pq"
)

// Postgres timeout
const Timeout = 5

// PgDB is a shortcut structure to a Postgres DB
type PgDB struct {
	*pg.DB
}

// Dial creates new database connection to postgres
func Dial(cfg *config.Config) (*PgDB, error) {
	pgOpts := &pg.Options{}

	var err error

	if cfg.PgURL != "" {
		pgOpts, err = pg.ParseURL(cfg.PgURL)
		if err != nil {
			return nil, err
		}
	}
	if cfg.PgProto != "" {
		pgOpts.Network = cfg.PgProto
	}
	if cfg.PgAddr != "" {
		pgOpts.Addr = cfg.PgAddr
	}
	if cfg.PgDb != "" {
		pgOpts.Database = cfg.PgDb
	}
	if cfg.PgUser != "" {
		pgOpts.User = cfg.PgUser
	}
	if cfg.PgPassword != "" {
		pgOpts.Password = cfg.PgPassword
	}

	pgDB := pg.Connect(pgOpts)

	_, err = pgDB.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	if Timeout > 0 {
		pgDB.WithTimeout(time.Second * time.Duration(Timeout))
	}

	return &PgDB{pgDB}, nil
}

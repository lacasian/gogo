package db

import (
	"database/sql"

	_ "github.com/kwix/gogo/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("module", "db")

type Config struct {
	ConnectionString string
	Automigrate      bool
}

type DB struct {
	config Config

	connection *sql.DB
}

func New(config Config) (*DB, error) {
	conn, err := sql.Open("postgres", config.ConnectionString)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if config.Automigrate {
		log.Info("attempting automatic execution of migrations")
		err = goose.Up(conn, "/")
		if err != nil && err != goose.ErrNoNextVersion {
			log.Fatal(err)
		}
		log.Info("database version is up to date")
	}

	return &DB{
		config:     config,
		connection: conn,
	}, nil
}

func (db *DB) Close() error {
	log.Info("Got stop signal")
	return db.connection.Close()
}

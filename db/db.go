package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

// common/global instance for singleton
var db *DB

func (db *DB) Close() {
	defer db.Pool.Close()
}

func ConnectDB() (*DB, error) {
	// singleton check
	if db != nil {
		return db, nil
	}

	config, err := pgxpool.ParseConfig("")
	config.ConnConfig.Host = "localhost"
	config.ConnConfig.Port = 5432
	config.ConnConfig.Database = "wishstore_db"

	if err != nil {
		return nil, fmt.Errorf("db error: %v", err)
	}
	dbPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("db error: %v", err)
	}
	db := DB{Pool: dbPool}
	return &db, nil
}

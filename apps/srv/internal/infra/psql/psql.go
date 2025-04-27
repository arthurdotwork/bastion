package psql

import (
	"context"
	"fmt"
	"time"

	"github.com/heetch/sqalx"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Queryable = sqalx.Node

type DB struct {
	Queryable

	db *sqlx.DB
}

func Connect(ctx context.Context, username string, password string, host string, port string, dbname string) (DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbname)

	db, err := sqlx.ConnectContext(ctx, "pgx", dsn)
	if err != nil {
		return DB{}, fmt.Errorf("failed to connect: %w", err)
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	if err := db.PingContext(ctx); err != nil {
		return DB{}, fmt.Errorf("failed to ping database: %w", err)
	}

	node, err := sqalx.New(db)
	if err != nil {
		return DB{}, fmt.Errorf("failed to create sqalx node: %w", err)
	}

	return DB{node, db}, nil
}

func (db *DB) Close() error {
	if err := db.Queryable.Close(); err != nil {
		return fmt.Errorf("failed to close sqalx node: %w", err)
	}

	if err := db.db.Close(); err != nil {
		return fmt.Errorf("failed to close database: %w", err)
	}

	return nil
}

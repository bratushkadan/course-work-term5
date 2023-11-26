package db

import (
	"context"
	"time"

	"floral/generated/database"

	"github.com/jackc/pgx/v5/pgxpool"

	"floral/config"

	_ "github.com/lib/pq"
)

// TODO: create pool of connections instead of opening a connection on earch

type User struct {
	Id       int
	Username string
}

var db *pgxpool.Pool

func init() {
	cfg, err := pgxpool.ParseConfig(config.App.Postgres.Connstr)
	if err != nil {
		panic(err)
	}

	cfg.MaxConnLifetime = 3 * time.Minute
	cfg.MinConns = 3
	cfg.MaxConns = 10
	cfg.HealthCheckPeriod = 30 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	pool, err := pgxpool.NewWithConfig(ctx, cfg)

	if err != nil {
		panic(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = pool.Ping(ctx)
	if err != nil {
		panic(err)
	}

	db = pool
}

func NewQueries() *database.Queries {
	return database.New(db)
}

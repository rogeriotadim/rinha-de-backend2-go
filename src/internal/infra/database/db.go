package database

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rogeriotadim/rinha-de-backend2-go/cmd/config"
)

var (
	db   *pgxpool.Pool
	once sync.Once
)

func NewDatabasePool(conf *config.Conf, ctx context.Context) (*pgxpool.Pool, error) {
	var err error = nil

	once.Do(func() {
		connUrl := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			conf.DBUser,
			conf.DBPassword,
			conf.DBHost,
			conf.DBPort,
			conf.DBName,
		)

		poolConfig, err := pgxpool.ParseConfig(connUrl)
		if err != nil {
			log.Fatalln("Unable to parse connection url:", err)
		}
		db, err = pgxpool.NewWithConfig(ctx, poolConfig)
		if err != nil {
			log.Fatalln("Unable to create connection pool:", err)
		}

		if err := db.Ping(ctx); err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
	})

	return db, err
}

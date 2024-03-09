package database

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	db   *pgxpool.Pool
	once sync.Once
	ctx  context.Context
)

func Connect(user string, password string, host string, port string, name string, max string) *pgxpool.Pool {
	once.Do(func() {
		url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, name)
		ctx = context.Background()

		maxConns, err := strconv.ParseInt(max, 10, 32)
		if err != nil {
			log.Fatalln("Unable to parse maximum pool size:", err)
		}

		for i := 0; i < 10; i++ {
			pool, err := pgxpool.ParseConfig(url)
			if err != nil {
				log.Fatalln("Unable to parse connection url:", err)
			}

			pool.MaxConns = int32(maxConns)

			db, err = pgxpool.NewWithConfig(ctx, pool)
			if err != nil {
				log.Fatalln("Unable to create connection pool:", err)
			}

			if err := db.Ping(ctx); err == nil {
				break
			} else {
				fmt.Println("Failed to connect to DB, retrying in 3 seconds")
				time.Sleep(3 * time.Second)
			}
		}
	})

	return db
}

func Exec(sql string) error {
	_, err := db.Exec(ctx, sql)
	return err
}

func Row(sql string, args ...any) pgx.Row {
	return db.QueryRow(ctx, sql, args...)
}

func Rows(sql string, args ...any) (pgx.Rows, error) {
	return db.Query(ctx, sql, args...)
}

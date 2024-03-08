package database

import (
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	db   *pgxpool.Pool
	once sync.Once
	ctx  context.Context
)

func Connect(url string) *pgxpool.Pool {
	once.Do(func() {
		ctx = context.Background()

		pool, err := pgxpool.ParseConfig(url)
		if err != nil {
			log.Fatalln("Unable to parse connection url:", err)
		}

		db, err = pgxpool.NewWithConfig(ctx, pool)
		if err != nil {
			log.Fatalln("Unable to create connection pool:", err)
		}

		if err := db.Ping(ctx); err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		truncate := `TRUNCATE TABLE public.client RESTART IDENTITY;
		TRUNCATE TABLE public.client_transaction RESTART IDENTITY;
		INSERT INTO public.client(id, "limit", balance) VALUES (1, 100000, 0);
		INSERT INTO public.client(id, "limit", balance) VALUES (2, 80000, 0);
		INSERT INTO public.client(id, "limit", balance) VALUES (3, 1000000, 0);
		INSERT INTO public.client(id, "limit", balance) VALUES (4, 10000000, 0);
		INSERT INTO public.client(id, "limit", balance) VALUES (5, 500000, 0);`
		if err := Exec(truncate); err != nil {
			log.Fatalf("Failed to truncante: %v", err)
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

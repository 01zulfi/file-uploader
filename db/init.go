package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool
var ctx = context.Background()

func Init() error {
	databaseUrl, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	newPool, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error while connecting to database")
		return err
	}
	pool = newPool

	sql, err := os.ReadFile("./db/init.sql")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = pool.Exec(ctx, string(sql))

	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

func Get() *pgxpool.Pool {
	return pool
}

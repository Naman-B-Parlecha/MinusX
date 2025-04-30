package db

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDb() (*sql.DB, error) {
	err := godotenv.Load()
	dsn := os.Getenv("DSN")

	db, err := openDb(dsn)
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

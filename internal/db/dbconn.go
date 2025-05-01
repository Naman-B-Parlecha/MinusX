package db

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"time"

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

	err = DBTablesInit(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func DBTablesInit(db *sql.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `CREATE TABLE IF NOT EXISTS Users (ID VARCHAR(255) NOT NULL, Username VARCHAR(255) NOT NULL UNIQUE, Email VARCHAR(255) NOT NULL UNIQUE, Password VARCHAR(255)  NOT NULL, Created_at TIMESTAMP NOT NULL, Updated_at TIMESTAMP NOT NULL, Avatar VARCHAR(255) NOT NULL, PRIMARY KEY (ID));`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return errors.New("failed to create users table: " + err.Error())
	}

	query = `CREATE TABLE IF NOT EXISTS Posts (ID VARCHAR(255) NOT NULL, UserID VARCHAR(255) NOT NULL, Title VARCHAR(255) NOT NULL, Content TEXT NOT NULL, ImageURL VARCHAR(255), Views INTEGER, Created_at TIMESTAMP NOT NULL, Updated_at TIMESTAMP NOT NULL, PRIMARY KEY (ID), FOREIGN KEY(UserID) REFERENCES Users(ID));`

	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return errors.New("failed to create users table: " + err.Error())
	}

	query = `CREATE TABLE IF NOT EXISTS Comments (ID VARCHAR(255) NOT NULL, UserID VARCHAR(255) NOT NULL, PostID VARCHAR(255) NOT NULL, Content TEXT NOT NULL, Created_at TIMESTAMP NOT NULL, Updated_at TIMESTAMP NOT NULL, PRIMARY KEY (ID), FOREIGN KEY(UserID) REFERENCES Users(ID), FOREIGN KEY(PostID) REFERENCES Posts(ID));`
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return errors.New("failed to create users table: " + err.Error())
	}

	return nil
}

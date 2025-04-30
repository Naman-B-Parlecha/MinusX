package services

import "database/sql"

type AuthService struct {
	DB *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{
		DB: db,
	}
}

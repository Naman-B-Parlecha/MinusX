package services

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/Naman-B-Parlecha/MinusX/internal/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type AuthService struct {
	DB *sql.DB
}

type Claims struct {
	UserID     string    `json:"user_id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Avatar     string    `json:"avatar"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	jwt.RegisteredClaims
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{
		DB: db,
	}
}

func (s *AuthService) RegisterUser(username, email, password, avatar string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if username == "" || email == "" || password == "" {
		return errors.New("username, email, and password cannot be empty")
	}

	if avatar == "" {
		avatar = "https://i.pinimg.com/736x/a4/5d/6a/a45d6a33e6328b8c16763bfc1462f8fa.jpg"
	}

	generatedId := uuid.New()

	hashedPassword, err := util.GenerateHashPass(password)
	if err != nil {
		return errors.Errorf("failed to hash password: %v", err)
	}

	query := `INSERT INTO Users (ID, Username, Email, Password, Created_at, Updated_at, Avatar) VALUES ($1, $2, $3, $4, $5, $6, $7);`

	_, err = s.DB.ExecContext(ctx, query, generatedId, username, email, hashedPassword, time.Now().UTC(), time.Now().UTC(), avatar)

	if err != nil {
		return errors.Errorf("failed to register user: %v", err)
	}
	return nil
}

func (s *AuthService) LoginUser(email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if email == "" || password == "" {
		return "", errors.New("email and password cannot be empty")
	}

	query := `SELECT ID, Username, Password, Created_at, Updated_at FROM Users WHERE Email = $1;`
	var userID, username, hashedPassword string
	var created_at, updated_at time.Time
	err := s.DB.QueryRowContext(ctx, query, email).Scan(&userID, &username, &hashedPassword, &created_at, &updated_at)
	if err != nil {
		return "", errors.Errorf("failed to find user: %v", err)
	}

	err = util.VerifyPassword(hashedPassword, password)
	if err != nil {
		return "", errors.Errorf("failed to verify password: %v", err)
	}

	claims := &Claims{
		UserID:     userID,
		Username:   username,
		Email:      email,
		Created_at: created_at,
		Updated_at: updated_at,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	godotenv.Load()
	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.Errorf("failed to generate token: %v", err)
	}

	return tokenString, nil

}

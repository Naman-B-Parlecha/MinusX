package routes

import (
	"database/sql"

	"github.com/Naman-B-Parlecha/MinusX/internal/handlers"
	"github.com/Naman-B-Parlecha/MinusX/internal/services"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, db *sql.DB) {
	authService := services.NewAuthService(db)
	authHandler := handlers.NewAuthHandler(authService)

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
}

package routes

import (
	"database/sql"

	"github.com/Naman-B-Parlecha/MinusX/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.RouterGroup, db *sql.DB) {

	r.Use(middlewares.JWTMiddleware())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "heepp",
		})
	})
}

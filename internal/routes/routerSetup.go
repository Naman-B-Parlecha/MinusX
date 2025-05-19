package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to MinusX API",
		})
	})
	r.GET("health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "sax sux",
		})
	})
	authGroup := r.Group("/auth")
	{
		AuthRoutes(authGroup, db)
	}

	postGroup := r.Group("/posts")
	{
		PostRoutes(postGroup, db)
	}

	commentGroup := r.Group("/comments")
	{
		CommentRoutes(commentGroup, db)
	}
}

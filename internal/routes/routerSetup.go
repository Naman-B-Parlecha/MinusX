package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	authGroup := r.Group("/auth")
	{
		AuthRoutes(authGroup, db)
	}

	postGroup := r.Group("/posts")
	{
		PostRoutes(postGroup, db)
	}
}

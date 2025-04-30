package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {

	authGroup := r.Group("/auth")
	{
		AuthRoutes(authGroup)
	}
}

package routes

import (
	"database/sql"

	"github.com/Naman-B-Parlecha/MinusX/internal/handlers"
	"github.com/Naman-B-Parlecha/MinusX/internal/middlewares"
	"github.com/Naman-B-Parlecha/MinusX/internal/services"
	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.RouterGroup, db *sql.DB) {

	postService := services.NewPostService(db)
	postHandler := handlers.NewPostHandler(postService)

	r.GET("/", postHandler.GetAllPosts)
	r.GET("/:id", postHandler.GetPostByID)
	r.POST("/increment/:id", postHandler.IncrementPostViews)

	r.Use(middlewares.JWTMiddleware())

	r.POST("/create", postHandler.CreatePost)
	r.PUT("/:id", postHandler.UpdatePost)
	r.DELETE("/:id", postHandler.DeletePost)
}

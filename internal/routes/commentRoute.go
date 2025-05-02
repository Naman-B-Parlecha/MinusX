package routes

import (
	"database/sql"

	"github.com/Naman-B-Parlecha/MinusX/internal/handlers"
	"github.com/Naman-B-Parlecha/MinusX/internal/middlewares"
	"github.com/Naman-B-Parlecha/MinusX/internal/services"
	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.RouterGroup, db *sql.DB) {
	commentService := services.NewCommentService(db)
	commentHandler := handlers.NewCommentHandler(commentService)

	r.Use(middlewares.JWTMiddleware())
	r.POST("/:id", commentHandler.AddNewComment)
	r.PUT("/:id", commentHandler.UpdateComment)
	r.DELETE("/:id", commentHandler.DeleteComment)
}

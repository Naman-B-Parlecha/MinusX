package handlers

import (
	"net/http"

	"github.com/Naman-B-Parlecha/MinusX/internal/middlewares"
	"github.com/Naman-B-Parlecha/MinusX/internal/services"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService *services.CommentService
}

func NewCommentHandler(commentService *services.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

func (h *CommentHandler) AddNewComment(c *gin.Context) {
	var Comment struct {
		Content string `json:"content" binding:"required"`
	}
	postId := c.Param("id")

	if postId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "post not found",
		})
		return
	}

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized to make post",
		})
		return
	}

	claim, ok := claims.(*middlewares.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized to make post",
		})
		return
	}

	if err := c.ShouldBindJSON(&Comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid payload",
		})
		return
	}

	resp, err := h.commentService.AddNewComment(postId, claim.UserID, Comment.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to add comment",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment Added Successfully",
		"data":    resp,
	})
}

func (h *CommentHandler) UpdateComment(c *gin.Context) {
	var Comment struct {
		Content string `json:"content" binding:"required"`
	}
	commentId := c.Param("id")

	if commentId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "comment not found",
		})
		return
	}

	existingComment, err := h.commentService.GetcommentByID(commentId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "comment not found",
		})
		return
	}

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized to make post",
		})
		return
	}

	claim, ok := claims.(*middlewares.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized to make post",
		})
		return
	}

	if err := c.ShouldBindJSON(&Comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid payload",
		})
		return
	}

	if existingComment.UserID != claim.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized to make post",
		})
		return
	}

	resp, err := h.commentService.UpdateComment(commentId, Comment.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update comment",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment updated Successfully",
		"data":    resp,
	})
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	commentId := c.Param("id")

	if commentId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "comment not found",
		})
		return
	}

	existingComment, err := h.commentService.GetcommentByID(commentId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "comment not found",
		})
		return
	}

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized to make post",
		})
		return
	}

	claim, ok := claims.(*middlewares.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized to make post",
		})
		return
	}

	if existingComment.UserID != claim.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized to make post",
		})
		return
	}

	err = h.commentService.DeleteComment(commentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update comment",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment Deleted Successfully",
	})
}

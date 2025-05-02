package handlers

import (
	"net/http"

	"github.com/Naman-B-Parlecha/MinusX/internal/middlewares"
	"github.com/Naman-B-Parlecha/MinusX/internal/services"
	"github.com/gin-gonic/gin"
)

type PostHander struct {
	postService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHander {
	return &PostHander{
		postService: postService,
	}
}

func (h *PostHander) GetAllPosts(c *gin.Context) {
	posts, err := h.postService.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Posts fetched successfully", "data": posts})
}

func (h *PostHander) CreatePost(c *gin.Context) {
	var Post struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		Image   string `json:"image" binding:"required"`
	}

	if err := c.ShouldBindJSON(&Post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": err.Error()})
		return
	}
	claimsValue, exists := c.Get("claims")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	claims, ok := claimsValue.(*middlewares.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	resp, err := h.postService.CreatePost(Post.Title, Post.Content, claims.UserID, Post.Image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully", "data": gin.H{
		"id": resp.ID,
		"user": gin.H{
			"id":       claims.UserID,
			"username": claims.Username,
			"email":    claims.Email,
			"avatar":   claims.Avatar,
		},
		"title":     resp.Title,
		"content":   resp.Content,
		"image":     resp.ImageUrl,
		"views":     resp.Views,
		"createdAt": resp.CreatedAt,
		"updatedAt": resp.UpdatedAt,
	}})
}

func (h *PostHander) GetPostByID(c *gin.Context) {
	postID := c.Param("id")
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
		return
	}

	post, err := h.postService.GetPostByID(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}

	if post == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post fetched successfully", "data": gin.H{
		"id":        post.ID,
		"userid":    post.UserID,
		"title":     post.Title,
		"content":   post.Content,
		"image":     post.ImageUrl,
		"views":     post.Views,
		"createdAt": post.CreatedAt,
		"updatedAt": post.UpdatedAt,
	}})
}

func (h *PostHander) UpdatePost(c *gin.Context) {
	postID := c.Param("id")
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
		return
	}

	var Post struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Image   string `json:"image"`
	}

	if err := c.ShouldBindJSON(&Post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "message": err.Error()})
		return
	}
	existingPost, err := h.postService.GetPostByID(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	if existingPost == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	title := Post.Title
	if title == "" {
		title = existingPost.Title
	}

	content := Post.Content
	if content == "" {
		content = existingPost.Content
	}

	image := Post.Image
	if image == "" {
		image = existingPost.ImageUrl
	}

	claimsValue, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	claims, ok := claimsValue.(*middlewares.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if claims.UserID != existingPost.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this post"})
		return
	}

	resp, err := h.postService.UpdatePost(postID, title, content, image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully", "data": resp})
}

func (h *PostHander) DeletePost(c *gin.Context) {
	postID := c.Param("id")
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
		return
	}

	claimsValue, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	claims, ok := claimsValue.(*middlewares.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	existingPost, err := h.postService.GetPostByID(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	if existingPost == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if claims.UserID != existingPost.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this post"})
		return
	}

	err = h.postService.DeletePost(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func (h *PostHander) IncrementPostViews(c *gin.Context) {
	postID := c.Param("id")
	if postID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post ID is required"})
		return
	}

	err := h.postService.IncrementPostViews(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post views incremented successfully"})
}

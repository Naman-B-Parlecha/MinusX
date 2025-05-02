package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/Naman-B-Parlecha/MinusX/models"
	"github.com/google/uuid"
)

type CommentService struct {
	db *sql.DB
}

func NewCommentService(db *sql.DB) *CommentService {
	return &CommentService{
		db: db,
	}
}

func (cs *CommentService) AddNewComment(postID string, userID string, content string) (*models.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	comment := &models.Comment{}

	id := uuid.New().String()
	query := "INSERT INTO Comments (id, postid, userid, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, postid, userid, content, created_at, updated_at"
	err := cs.db.QueryRowContext(ctx, query, id, postID, userID, content, time.Now(), time.Now()).Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return &models.Comment{}, err
	}
	return comment, nil
}

func (cs *CommentService) UpdateComment(commentID string, content string) (*models.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	comment := &models.Comment{}

	query := "UPDATE Comments SET content = $1, updated_at = $2 WHERE id = $3 RETURNING id, postid, userid, content, created_at, updated_at"
	err := cs.db.QueryRowContext(ctx, query, content, time.Now(), commentID).Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return &models.Comment{}, err
	}
	return comment, nil
}

func (cs *CommentService) DeleteComment(commentID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := "DELETE FROM Comments WHERE id = $1"
	_, err := cs.db.ExecContext(ctx, query, commentID)
	if err != nil {
		return err
	}
	return nil
}

func (s *CommentService) GetcommentByID(id string) (*models.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	comment := &models.Comment{}
	query := `SELECT id, userid, postid, content, created_at, updated_at FROM Comments WHERE id = $1`

	err := s.db.QueryRowContext(ctx, query, id).Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return comment, nil
}

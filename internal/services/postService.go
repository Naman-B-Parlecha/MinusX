package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/Naman-B-Parlecha/MinusX/models"
	"github.com/google/uuid"
)

type PostService struct {
	db *sql.DB
}

func NewPostService(db *sql.DB) *PostService {
	return &PostService{
		db: db,
	}
}

func (s *PostService) CreatePost(title, content, userID, imageURL string) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `INSERT INTO Posts (id, userid, title, content, imageurl, views, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, 0, $6, $7) RETURNING id, userid, title, content, imageurl, views, created_at, updated_at`

	id := uuid.New().String()
	now := time.Now()

	post := &models.Post{}
	err := s.db.QueryRowContext(ctx, query, id, userID, title, content, imageURL, now, now).Scan(
		&post.ID, &post.UserID, &post.Title, &post.Content, &post.ImageUrl, &post.Views, &post.CreatedAt, &post.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostService) GetAllPosts() ([]*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `SELECT id, userid, title, content, imageurl, views, created_at, updated_at FROM Posts`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		post := &models.Post{}
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.ImageUrl, &post.Views, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (s *PostService) GetPostByID(id string) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `SELECT id, userid, title, content, imageurl, views, created_at, updated_at FROM Posts WHERE id = $1`
	post := &models.Post{}

	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID, &post.UserID, &post.Title, &post.Content, &post.ImageUrl, &post.Views, &post.CreatedAt, &post.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return post, nil
}

func (s *PostService) UpdatePost(id, title, content, imageURL string) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `UPDATE Posts SET title = $1, content = $2, imageurl = $3, updated_at = $4 WHERE id = $5 RETURNING id, userid, title, content, imageurl, views, created_at, updated_at`
	post := &models.Post{}
	err := s.db.QueryRowContext(ctx, query, title, content, imageURL, time.Now(), id).Scan(
		&post.ID, &post.UserID, &post.Title, &post.Content, &post.ImageUrl, &post.Views, &post.CreatedAt, &post.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return post, nil
}

func (s *PostService) DeletePost(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `DELETE FROM Posts WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) IncrementPostViews(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `UPDATE Posts SET views = views + 1 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/mbasim25/go-http-server"
)

type CommentStore struct {
	*sqlx.DB
}

func (s *CommentStore) Comment(id int64) (posts.Comment, error) {
	var c posts.Comment
	if err := s.Get(&c, `SELECT * FROM comments WHERE id = $1`, id); err != nil {
		return posts.Comment{}, fmt.Errorf("error getting comment: %w", err)
	}

	return c, nil
}

func (s *CommentStore) CommentsByPost(postId int64) ([]posts.Comment, error) {
	var cc []posts.Comment
	if err := s.Select(&cc, `SELECT * FROM comments WHERE post_id = $1`, postId); err != nil {
		return []posts.Comment{}, fmt.Errorf("error getting comments: %w", err)
	}

	return cc, nil
}

func (s *CommentStore) CreateComment(c *posts.Comment) error {
	if err := s.Get(c, `INSERT INTO comments (post_id, body) VALUES($1, $2) RETURNING *`,
		c.PostId,
		c.Body); err != nil {
		return fmt.Errorf("error getting comments: %w", err)
	}

	return nil
}

func (s *CommentStore) UpdateComment(c *posts.Comment) error {
	if err := s.Get(c, `UPDATE comments SET body = $1 WHERE id = $2 RETURNING *`,
		c.Body,
		c.ID); err != nil {
		return fmt.Errorf("error getting comments: %w", err)
	}

	return nil
}

func (s *CommentStore) DeleteComment(id int64) error {
	if _, err := s.Exec(`DELETE * FROM comments WHERE id = $2`); err != nil {
		return fmt.Errorf("error getting comments: %w", err)
	}

	return nil
}

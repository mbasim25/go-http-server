package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/mbasim25/go-http-server"
)

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id int64) (posts.Post, error) {
	var p posts.Post

	if err := s.Get(&p, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return posts.Post{}, fmt.Errorf("error: %w", err)
	}

	return p, nil
}

func (s *PostStore) Posts() ([]posts.Post, error) {
	var pp []posts.Post

	if err := s.Select(&pp, `SELECT * FROM posts`); err != nil {
		return []posts.Post{}, fmt.Errorf("error getting posts: %w", err)
	}

	return pp, nil
}

func (s *PostStore) CreatePost(p *posts.Post) error {
	if _, err := s.NamedExec(`INSERT INTO posts (content) VALUES (:content) RETURNING *`,
		map[string]interface{}{
			"content": p.Content,
		}); err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}

	return nil
}

func (s *PostStore) UpdatePost(p *posts.Post) error {
	if err := s.Get(p, `UPDATE posts SET content = $1 WHERE id = $2 RETURNING *`,
		p.Content,
		p.ID); err != nil {
		return fmt.Errorf("error updating post: %w", err)
	}

	return nil
}

func (s *PostStore) DeletePost(id int64) error {
	if _, err := s.Exec(`DELETE FROM posts WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}

	return nil
}

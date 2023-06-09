package main

type Post struct {
	ID      int    `db:"id"`
	content string `db:"content"`
}

type Comment struct {
	ID      int    `db:"id"`
	post_id int    `db:"post_id"`
	body    string `db:"body"`
}

type PostStore interface {
	Post(id int) (Post, error)
	Posts() ([]Post, error)
	CreatePost(p *Post) error
	UpdatePost(p *Post) error
	DeletePost(id int) error
}

type CommentStore interface {
	Comment(id int) (Comment, error)
	CommentsByPost(post_id int) ([]Comment, error)
	CreateComment(c *Comment) error
	UpdateComment(c *Comment) error
	DeleteComment(id int) error
}

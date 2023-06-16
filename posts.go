package posts

type Post struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
}

type Comment struct {
	ID     int64  `db:"id"`
	PostId int64  `db:"post_id"`
	Body   string `db:"body"`
}

type PostStore interface {
	Post(id int64) (Post, error)
	Posts() ([]Post, error)
	CreatePost(p *Post) error
	UpdatePost(p *Post) error
	DeletePost(id int64) error
}

type CommentStore interface {
	Comment(id int64) (Comment, error)
	CommentsByPost(postId int64) ([]Comment, error)
	CreateComment(c *Comment) error
	UpdateComment(c *Comment) error
	DeleteComment(id int64) error
}

type Store interface {
	PostStore
	CommentStore
}

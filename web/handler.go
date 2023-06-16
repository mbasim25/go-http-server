package web

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/mbasim25/go-http-server"
)

func NewHandler(store posts.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)

	h.Route("/posts", func(r chi.Router) {
		r.Get("/", h.PostsList())
		r.Get("/new", h.PostsCreate())
		r.Post("/", h.PostsStore())
	}) // TODO: finish adding other routes/handlers
	return h
}

type Handler struct {
	*chi.Mux
	store posts.Store
}

const postsListHTML = `
  <h1>Posts</h1>
  <dl>
  {{range .Posts}}  
    <dt><strong>{{.ID}}</strong></dt>
    <dd>{{.Content}}</dd> 
  {{end}}
  </dl>
  `

func (h *Handler) PostsList() http.HandlerFunc {
	type data struct {
		Posts []posts.Post
	}

	tmpl := template.Must(template.New("").Parse(postsListHTML))

	return func(w http.ResponseWriter, _ *http.Request) {
		tt, err := h.store.Posts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data{Posts: tt})
	}
}

const postCreateHTML = `
  <h1>New Post</h1>
  <form action="/posts" method="POST">
    <table>
      <tr>
        <td>Content</td>
        <td><input type="text" name="content" /></td>
      </tr>
    </table>
    <button type="submit">Create Post</button>
  </form>
  `

func (h *Handler) PostsCreate() http.HandlerFunc {
	tmpl := template.Must(template.New("").Parse(postCreateHTML))
	return func(w http.ResponseWriter, _ *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func (h *Handler) PostsStore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := r.FormValue("content")

		if err := h.store.CreatePost(&posts.Post{
			Content: content,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/posts", http.StatusFound)
	}
}

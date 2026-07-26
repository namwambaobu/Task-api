package router

import (
	"net/http"

	"github.com/namwamba/task-api/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func New() http.Handler {
	r := chi.NewRouter()

	r.Get("/tasks", handlers.GetTasks)
	return r
}

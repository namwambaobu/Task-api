package router

import (
	"net/http"

	"github.com/namwamba/task-api/internal/handlers"
	"github.com/namwamba/task-api/internal/store"

	"github.com/go-chi/chi/v5"
)

func New() http.Handler {
	r := chi.NewRouter()

	taskStore := store.NewMemoryStore()
	taskHandler := handlers.NewTaskHandler(taskStore)

	r.Get("/tasks", taskHandler.GetTasks)
	return r
}

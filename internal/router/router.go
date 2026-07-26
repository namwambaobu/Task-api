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

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", taskHandler.GetTasks)
		r.Post("/", taskHandler.CreateTask)
		r.Get("/{id}", taskHandler.GetTask)
		r.Put("/{id}", taskHandler.UpdateTask)
		r.Delete("/{id}", taskHandler.DeleteTask)
	})
	return r
}

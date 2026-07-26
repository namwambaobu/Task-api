package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/namwamba/task-api/internal/store"
)

type TaskHandler struct {
	store store.TaskStore
}

func NewTaskHandler(store store.TaskStore) *TaskHandler {
	return &TaskHandler{
		store: store,
	}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	task := h.store.GetAll()

	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single task
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	task, found := h.store.GetByID(id)
	if !found {
		writeError(w, http.StatusNotFound, "Task not found")
		return
	}

	writeJSON(w, http.StatusOK, task)

}

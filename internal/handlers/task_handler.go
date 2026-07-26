package handlers

import (
	"encoding/json"
	"net/http"

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

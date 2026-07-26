package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/namwamba/task-api/internal/models"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {

	tasks := []models.Task{
		{
			ID:    1,
			Title: "Learn Go",
			Done:  false,
		},
		{
			ID:    2,
			Title: "Learn Chi",
			Done:  false,
		},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)
}

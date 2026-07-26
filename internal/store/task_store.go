package store

import "github.com/namwamba/task-api/internal/models"

type TaskStore interface {
	GetAll() []models.Task
	GetByID(id int) (models.Task, bool)
	Create(task models.Task) models.Task
	Update(id int, task models.Task) (models.Task, bool)
	Delete(id int) bool
}

package store

import "github.com/namwamba/task-api/internal/models"

type MemoryStore struct {
	tasks  []models.Task
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks: []models.Task{
			{ID: 1, Title: "Learn Go", Done: false},
			{ID: 2, Title: "Learn Chi", Done: false},
		},
		nextID: 3,
	}
}

func (s *MemoryStore) GetAll() []models.Task {
	return s.tasks
}

func (s *MemoryStore) GetByID(id int) (models.Task, bool) {
	for _, task := range s.tasks {
		if task.ID == id {
			return task, true
		}
	}
	return models.Task{}, false
}

func (s *MemoryStore) Create(task models.Task) models.Task {
	task.ID = s.nextID
	s.nextID++
	s.tasks = append(s.tasks, task)
	return task
}

func (s *MemoryStore) Update(id int, updatedTask models.Task) (models.Task, bool) {
	for i, task := range s.tasks {
		if task.ID == id {
			updatedTask.ID = id
			s.tasks[i] = updatedTask
			return updatedTask, true
		}
	}
	return models.Task{}, false
}

func (s *MemoryStore) Delete(id int) bool {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return true
		}
	}
	return false
}

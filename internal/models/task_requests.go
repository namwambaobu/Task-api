package models

type CreateTaskRequest struct {
	Title string `json:"title"`
}

type UpdateTaskRequest struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

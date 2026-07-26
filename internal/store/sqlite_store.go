package store

import (
	"context"
	"database/sql"

	"github.com/namwamba/task-api/internal/models"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(db *sql.DB) *SQLiteStore {
	return &SQLiteStore{
		db: db,
	}
}

func (s *SQLiteStore) GetAll(ctx context.Context) ([]models.Task, error) {

	rows, err := s.db.QueryContext(
		ctx,
		`SELECT id, title, done
		 FROM tasks
		 ORDER BY id`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {

		var task models.Task

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Done,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

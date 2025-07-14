package tasks

import "database/sql"

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	query := `SELECT id, description, completed FROM tasks`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Description, &task.IsCompleted); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

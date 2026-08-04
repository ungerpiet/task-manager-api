package repository

import (
	"context"
	"go-task-api/internal/model"

	"github.com/jackc/pgx/v5"
)

type TaskRepository struct {
	db *pgx.Conn
}

func NewTaskRepository(db *pgx.Conn) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAll() ([]model.Task, error) {
	rows, err := r.db.Query(context.Background(), "SELECT id, title, description, completed, created_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var t model.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) Create(task *model.Task) error {
	err := r.db.QueryRow(
		context.Background(),
		"INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id, completed, created_at",
		task.Title,
		task.Description,
	).Scan(&task.ID, &task.Completed, &task.CreatedAt)

	return err
}

func (r *TaskRepository) GetByID(id int) (*model.Task, error) {
	var t model.Task
	err := r.db.QueryRow(
		context.Background(),
		"SELECT id, title, description, completed, created_at FROM tasks WHERE id = $1",
		id,
	).Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *TaskRepository) Update(task *model.Task, id int) error {
	commandTag, err := r.db.Exec(
		context.Background(),
		"UPDATE tasks SET title = $1, description = $2, completed = $3 WHERE id = $4",
		task.Title,
		task.Description,
		task.Completed,
		id,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

//func (r *TaskRepository) Create(task *model.Task) error {}

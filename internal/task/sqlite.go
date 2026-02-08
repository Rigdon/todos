package task

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) Migrate(ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		status TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := r.db.ExecContext(ctx, query)
	return err
}

func (r *SQLiteRepository) Create(ctx context.Context, task Task) error {
	query := `INSERT INTO tasks (id, title, status, created_at) VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, task.ID, task.Title, task.Status, task.CreatedAt)
	return err
}

func (r *SQLiteRepository) Get(ctx context.Context, id string) (Task, error) {
	query := `SELECT id, title, status, created_at FROM tasks WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)

	var task Task
	err := row.Scan(&task.ID, &task.Title, &task.Status, &task.CreatedAt)
	if err == sql.ErrNoRows {
		return Task{}, ErrNotFound
	}
	return task, err
}

func (r *SQLiteRepository) List(ctx context.Context) ([]Task, error) {
	query := `SELECT id, title, status, created_at FROM tasks ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Status, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *SQLiteRepository) Update(ctx context.Context, id string, title, status *string) (Task, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return Task{}, err
	}
	defer tx.Rollback()

	// Check if exists
	var current Task
	queryGet := `SELECT id, title, status, created_at FROM tasks WHERE id = ?`
	err = tx.QueryRowContext(ctx, queryGet, id).Scan(&current.ID, &current.Title, &current.Status, &current.CreatedAt)
	if err == sql.ErrNoRows {
		return Task{}, ErrNotFound
	}
	if err != nil {
		return Task{}, err
	}

	newTitle := current.Title
	if title != nil {
		newTitle = *title
	}
	newStatus := current.Status
	if status != nil {
		newStatus = *status
	}

	queryUpdate := `UPDATE tasks SET title = ?, status = ? WHERE id = ?`
	_, err = tx.ExecContext(ctx, queryUpdate, newTitle, newStatus, id)
	if err != nil {
		return Task{}, err
	}

	if err := tx.Commit(); err != nil {
		return Task{}, err
	}

	current.Title = newTitle
	current.Status = newStatus
	return current, nil
}

func (r *SQLiteRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM tasks WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrNotFound
	}
	return nil
}

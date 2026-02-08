package task

import (
	"context"
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("task not found")
)

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Repository interface {
	Create(ctx context.Context, task Task) error
	Get(ctx context.Context, id string) (Task, error)
	List(ctx context.Context) ([]Task, error)
	Update(ctx context.Context, id string, title, status *string) (Task, error)
	Delete(ctx context.Context, id string) error
}

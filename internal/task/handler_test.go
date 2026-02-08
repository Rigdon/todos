package task

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type mockRepo struct {
	createFunc func(context.Context, Task) error
	getFunc    func(context.Context, string) (Task, error)
	listFunc   func(context.Context) ([]Task, error)
	updateFunc func(context.Context, string, *string, *string) (Task, error)
	deleteFunc func(context.Context, string) error
}

func (m *mockRepo) Create(ctx context.Context, t Task) error         { return m.createFunc(ctx, t) }
func (m *mockRepo) Get(ctx context.Context, id string) (Task, error) { return m.getFunc(ctx, id) }
func (m *mockRepo) List(ctx context.Context) ([]Task, error)         { return m.listFunc(ctx) }
func (m *mockRepo) Update(ctx context.Context, id string, title, status *string) (Task, error) {
	return m.updateFunc(ctx, id, title, status)
}
func (m *mockRepo) Delete(ctx context.Context, id string) error { return m.deleteFunc(ctx, id) }

func TestHandler_List(t *testing.T) {
	now := time.Now().UTC()
	tasks := []Task{
		{ID: "1", Title: "Task 1", Status: "todo", CreatedAt: now},
	}
	repo := &mockRepo{
		listFunc: func(ctx context.Context) ([]Task, error) {
			return tasks, nil
		},
	}
	h := NewHandler(repo)

	req := httptest.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()

	h.list(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var resp []Task
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(resp) != 1 {
		t.Errorf("expected 1 task, got %d", len(resp))
	}
}

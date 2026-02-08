package task

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestHandler_Create(t *testing.T) {
	repo := &mockRepo{
		createFunc: func(ctx context.Context, task Task) error {
			if task.Title != "New Task" {
				t.Errorf("expected title 'New Task', got '%s'", task.Title)
			}
			return nil
		},
	}
	h := NewHandler(repo)

	body := `{"title": "New Task"}`
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	h.create(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", w.Code)
	}

	var resp Task
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if resp.Title != "New Task" {
		t.Errorf("expected title 'New Task', got '%s'", resp.Title)
	}
	if resp.ID == "" {
		t.Error("expected ID to be generated")
	}
}

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

func TestHandler_Update(t *testing.T) {
	repo := &mockRepo{
		updateFunc: func(ctx context.Context, id string, title, status *string) (Task, error) {
			if id != "1" {
				t.Errorf("expected ID '1', got '%s'", id)
			}
			if *title != "Updated" {
				t.Errorf("expected title 'Updated', got '%s'", *title)
			}
			return Task{ID: "1", Title: "Updated", Status: "todo"}, nil
		},
	}
	h := NewHandler(repo)

	body := `{"title": "Updated"}`
	req := httptest.NewRequest("PUT", "/tasks/1", bytes.NewBufferString(body))
	req.SetPathValue("id", "1")
	w := httptest.NewRecorder()

	h.update(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestHandler_Delete(t *testing.T) {
	repo := &mockRepo{
		deleteFunc: func(ctx context.Context, id string) error {
			if id != "1" {
				t.Errorf("expected ID '1', got '%s'", id)
			}
			return nil
		},
	}
	h := NewHandler(repo)

	req := httptest.NewRequest("DELETE", "/tasks/1", nil)
	req.SetPathValue("id", "1")
	w := httptest.NewRecorder()

	h.delete(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected status 204, got %d", w.Code)
	}
}

func TestHandler_Create_BadRequest(t *testing.T) {
	h := NewHandler(&mockRepo{})

	// Empty body
	req := httptest.NewRequest("POST", "/tasks", strings.NewReader(""))
	w := httptest.NewRecorder()
	h.create(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 for empty body, got %d", w.Code)
	}

	// Missing title
	req = httptest.NewRequest("POST", "/tasks", strings.NewReader(`{}`))
	w = httptest.NewRecorder()
	h.create(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 for missing title, got %d", w.Code)
	}
}

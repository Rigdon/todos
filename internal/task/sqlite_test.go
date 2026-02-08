package task

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *SQLiteRepository {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	repo := NewSQLiteRepository(db)
	if err := repo.Migrate(context.Background()); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return repo
}

func TestSQLiteRepository_CRUD(t *testing.T) {
	repo := setupTestDB(t)
	ctx := context.Background()

	// 1. Create
	now := time.Now().UTC().Truncate(time.Second) // SQLite stores time with limited precision by default
	task := Task{
		ID:        "1",
		Title:     "Test Task",
		Status:    "todo",
		CreatedAt: now,
	}

	if err := repo.Create(ctx, task); err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	// 2. Get
	got, err := repo.Get(ctx, "1")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if got.ID != task.ID {
		t.Errorf("expected ID %s, got %s", task.ID, got.ID)
	}
	if got.Title != task.Title {
		t.Errorf("expected Title %s, got %s", task.Title, got.Title)
	}

	// 3. Update
	newTitle := "Updated Title"
	newStatus := "done"
	updated, err := repo.Update(ctx, "1", &newTitle, &newStatus)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	if updated.Title != newTitle {
		t.Errorf("expected Title %s, got %s", newTitle, updated.Title)
	}
	if updated.Status != newStatus {
		t.Errorf("expected Status %s, got %s", newStatus, updated.Status)
	}

	// 4. List
	list, err := repo.List(ctx)
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	if len(list) != 1 {
		t.Errorf("expected 1 task in list, got %d", len(list))
	}

	// 5. Delete
	if err := repo.Delete(ctx, "1"); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	_, err = repo.Get(ctx, "1")
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

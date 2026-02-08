# Task Management

## Purpose
Manage the lifecycle of tasks in the system.

## Requirements

### Requirement: Task Entity Schema
The system MUST store tasks with a defined schema.

#### Scenario: valid task structure
- **WHEN** a task is created or retrieved
- **THEN** it MUST contain an `id` (string/UUID)
- **AND** it MUST contain a `title` (string)
- **AND** it MUST contain a `status` (string: "todo", "in-progress", "done")
- **AND** it MUST contain `created_at` (timestamp)

### Requirement: Create Task
The system MUST allow creating new tasks via HTTP.

#### Scenario: create success
- **WHEN** a `POST` request is sent to `/tasks` with a JSON body `{"title": "Buy milk"}`
- **THEN** the system MUST return status `201 Created`
- **AND** the response body MUST include the created task with a generated `id` and default status "todo"

#### Scenario: create validation failure
- **WHEN** a `POST` request is sent to `/tasks` with an empty title
- **THEN** the system MUST return status `400 Bad Request`

### Requirement: List Tasks
The system MUST allow listing all tasks.

#### Scenario: list all
- **WHEN** a `GET` request is sent to `/tasks`
- **THEN** the system MUST return status `200 OK`
- **AND** the response body MUST be a JSON array of task objects

### Requirement: Update Task
The system MUST allow updating a task's status or title.

#### Scenario: update status
- **WHEN** a `PUT` request is sent to `/tasks/{id}` with `{"status": "done"}`
- **THEN** the system MUST return status `200 OK`
- **AND** the task status MUST be updated to "done"

#### Scenario: update non-existent
- **WHEN** a `PUT` request is sent to a non-existent ID
- **THEN** the system MUST return status `404 Not Found`

### Requirement: Delete Task
The system MUST allow deleting a task.

#### Scenario: delete success
- **WHEN** a `DELETE` request is sent to `/tasks/{id}`
- **THEN** the system MUST return status `204 No Content`
- **AND** subsequent `GET` requests for that ID MUST return `404 Not Found`

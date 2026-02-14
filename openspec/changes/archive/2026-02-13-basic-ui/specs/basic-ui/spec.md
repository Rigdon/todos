## ADDED Requirements

### Requirement: View Task List
The system SHALL display a list of all tasks fetched from the backend API.

#### Scenario: Display tasks
- **WHEN** the user opens the application
- **THEN** the system fetches tasks from `GET /tasks`
- **AND** displays each task's title and status

### Requirement: Create Task
The system SHALL allow users to create a new task.

#### Scenario: Add a new task
- **WHEN** the user enters a title in the input field
- **AND** clicks the "Add" button (or presses Enter)
- **THEN** the system sends a `POST /tasks` request with the title
- **AND** adds the new task to the displayed list

### Requirement: Update Task Status
The system SHALL allow users to toggle the status of a task between "todo" and "done".

#### Scenario: Toggle status
- **WHEN** the user clicks the checkbox next to a task
- **THEN** the system sends a `PUT /tasks/{id}` request with the new status
- **AND** updates the visual state of the task (e.g., strikethrough for done)

### Requirement: Edit Task Title
The system SHALL allow users to edit the title of an existing task.

#### Scenario: Edit title
- **WHEN** the user double-clicks the task title
- **AND** modifies the text and presses Enter
- **THEN** the system sends a `PUT /tasks/{id}` request with the new title
- **AND** updates the displayed title

### Requirement: Delete Task
The system SHALL allow users to delete a task.

#### Scenario: Delete a task
- **WHEN** the user clicks the "Delete" button next to a task
- **THEN** the system sends a `DELETE /tasks/{id}` request
- **AND** removes the task from the displayed list

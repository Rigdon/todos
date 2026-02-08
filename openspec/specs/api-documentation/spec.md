# API Documentation Spec

## Purpose
Provide a standardized OpenAPI specification for the API to ensure consistency and improve developer experience.

## Requirements

### Requirement: OpenAPI Specification
The system SHALL provide an OpenAPI 3.0 specification file documenting the API.

#### Scenario: Spec file exists
- **WHEN** checking the repository root
- **THEN** an `openapi.yaml` file exists

### Requirement: List Tasks Endpoint
The spec SHALL document the `GET /tasks` endpoint.

#### Scenario: List tasks definition
- **WHEN** inspecting `GET /tasks` in `openapi.yaml`
- **THEN** it describes a successful response with a list of tasks
- **AND** the task schema includes `id`, `title`, `status`, and `created_at`

### Requirement: Create Task Endpoint
The spec SHALL document the `POST /tasks` endpoint.

#### Scenario: Create task definition
- **WHEN** inspecting `POST /tasks` in `openapi.yaml`
- **THEN** it accepts a JSON body with a required `title` field
- **AND** it describes a 201 Created response with the created task

### Requirement: Update Task Endpoint
The spec SHALL document the `PUT /tasks/{id}` endpoint.

#### Scenario: Update task definition
- **WHEN** inspecting `PUT /tasks/{id}` in `openapi.yaml`
- **THEN** it accepts a JSON body with optional `title` and `status` fields
- **AND** it describes a 200 OK response with the updated task
- **AND** it describes a 404 Not Found response

### Requirement: Delete Task Endpoint
The spec SHALL document the `DELETE /tasks/{id}` endpoint.

#### Scenario: Delete task definition
- **WHEN** inspecting `DELETE /tasks/{id}` in `openapi.yaml`
- **THEN** it describes a 204 No Content response
- **AND** it describes a 404 Not Found response

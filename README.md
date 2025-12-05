Task Manager API (Go)

A simple concurrency-safe in-memory Task Manager API built in Go using Gorilla Mux.

Features
Create, read, update, and delete tasks (CRUD)
Concurrency-safe in-memory storage using sync.RWMutex
Input validation and JSON error handling
Simple RESTful API design

Requirements
Go 1.25+ installed
Optional: VS Code REST Client,Invoke-WebRequest or cURL, or Postman for testing

Setup & Run

Clone the repository:

git clone https://github.com/Thripada/task-manager-api-go.git
cd task-manager-api-go


Install dependencies:

go mod tidy


Run the server:

go run main.go


You should see:

Server running on port 8080


Project Structure

task-manager-api-go/
│── main.go
│── go.mod
│── go.sum
│── README.md
│── .gitignore
├── models/
│   └── task.go
├── internal/
│   ├── models/
|   |    └── task.go
│   └── store/
|       └── taskstore.go
├── handler/
|   └── tasks.go
└── tests/
    └── requests.rest


Endpoints



API: http://localhost:8080

Endpoints
1. POST /tasks create  (this command for powershell terminal)

Invoke-WebRequest -Uri "http://localhost:8080/tasks" `
  -Method POST `
  -Headers @{ "Content-Type"="application/json" } `
  -Body '{"title":"Learn Go","status":"pending"}'

2. List All Tasks   (this command for powershell terminal)

Invoke-WebRequest -Uri "http://localhost:8080/tasks" -Method GET

3. Get Task by ID (example id=1) 

Invoke-WebRequest -Uri "http://localhost:8080/tasks/1" -Method GET

4. PUT /tasks/{id} update  (example id=1)

Invoke-WebRequest -Uri "http://localhost:8080/tasks/1" `
  -Method PUT `
  -Headers @{ "Content-Type"="application/json" } `
  -Body '{"title":"Updated Go Task","status":"completed"}'


5. DELETE /tasks/{id} delete  (example id=1)

Invoke-WebRequest -Uri "http://localhost:8080/tasks/1" -Method DELETE

Test
Use the tests/requests.rest with the VS Code REST Client or use curl / PowerShell InvokeWebRequest .


Concurrency approach

The project uses an in-memory map to store tasks.
Since multiple requests can access or modify the map at the same time, a mutex (sync.RWMutex) is used.

RLock() is used when reading tasks

Lock() is used when adding, updating, or deleting tasks

This prevents data races and ensures the program works correctly even with multiple concurrent users.


Input Validation & Errors 

When creating a task, the title is required

If JSON is invalid → 400 Bad Request

If a task ID does not exist → 404 Not Found

Errors are returned in JSON format like:

{ "error": "task not found" }
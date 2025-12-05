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

git clone https://github.com/thripada/task-manager-api-go.git
cd task-manager-api-go


Install dependencies:

go mod tidy


Run the server:

go run main.go


You should see:

Server running on port 8080



API: http://localhost:8080

Endpoints
POST /tasks create
GET /tasks list
GET /tasks/{id} retrieve
PUT /tasks/{id} update
DELETE /tasks/{id} delete
Test
Use the tests/requests.rest with the VS Code REST Client or use curl / PowerShell InvokeWebRequest .


Concurrency approach
The in-memory store ( internal/store/taskstore.go ) uses sync.RWMutex to protect the
map.
Readers use RLock for concurrency; writers use Lock to ensure exclusive writes.



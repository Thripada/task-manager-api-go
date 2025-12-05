# Task Manager API (Go) — Final
## Setup
1. Replace module path
Open `go.mod` and `import` paths in source files and replace `github.com/
Thripada/task-manager-api-go` with your GitHub module path (e.g.
`github.com/Thripada/task-manager-api-go`).
2. Download dependencies
```bash
go mod tidy


1. Run locally

go run main.go


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


Publishing to GitHub
Update module path as above
Initialize git, commit, add remote, and push
git init
git add .
git commit -m "initial"
git branch -M main
git remote add origin https://github.com/<YOUR_USERNAME>/task-manager-apigo.git
git push -u origin main

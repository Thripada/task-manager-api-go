package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Thripada/task-manager-api-go/internal/models"
	"github.com/Thripada/task-manager-api-go/internal/store"
	"github.com/gorilla/mux"
)

type TasksHandler struct {
	Store *store.TaskStore
}

func NewTasksHandler(s *store.TaskStore) *TasksHandler { return &TasksHandler{Store: s} }

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

// Create POST /tasks
func (h *TasksHandler) Create(w http.ResponseWriter, r *http.Request) {
	var in models.CreateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	task, err := h.Store.Create(in)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, task)
}

// List GET /tasks
func (h *TasksHandler) List(w http.ResponseWriter, r *http.Request) {
	tasks := h.Store.List()
	writeJSON(w, http.StatusOK, tasks)
}

// Get GET /tasks/{id}
func (h *TasksHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	t, err := h.Store.Get(id)
	if err != nil {
		writeError(w, http.StatusNotFound, "task not found")
		return
	}
	writeJSON(w, http.StatusOK, t)
}

// Update PUT /tasks/{id}
func (h *TasksHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var in models.UpdateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	updated, err := h.Store.Update(id, in)
	if err != nil {
		writeError(w, http.StatusNotFound, "task not found")
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

// Delete DELETE /tasks/{id}
func (h *TasksHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.Store.Delete(id); err != nil {
		writeError(w, http.StatusNotFound, "task not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

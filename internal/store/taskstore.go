package store

import (
	"errors"
	"sync"
	"time"

	"github.com/Thripada/task-manager-api-go/internal/models"
	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("task not found")
)

type TaskStore struct {
	mu    sync.RWMutex
	store map[string]models.Task
}

func NewTaskStore() *TaskStore {
	return &TaskStore{store: make(map[string]models.Task)}
}

func (s *TaskStore) List() []models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]models.Task, 0, len(s.store))
	for _, v := range s.store {
		out = append(out, v)
	}
	return out
}

func (s *TaskStore) Get(id string) (models.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.store[id]
	if !ok {
		return models.Task{}, ErrNotFound
	}
	return t, nil
}

func (s *TaskStore) Create(input models.CreateTaskInput) (models.Task, error) {
	if input.Title == "" {
		return models.Task{}, errors.New("title is required")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.NewString()
	now := time.Now().UTC()
	completed := false
	if input.Completed != nil {
		completed = *input.Completed
	}
	t := models.Task{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
		Completed:   completed,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	s.store[id] = t
	return t, nil
}

func (s *TaskStore) Update(id string, input models.UpdateTaskInput) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	existing, ok := s.store[id]
	if !ok {
		return models.Task{}, ErrNotFound
	}
	if input.Title != nil {
		existing.Title = *input.Title
	}
	if input.Description != nil {
		existing.Description = *input.Description
	}
	if input.Completed != nil {
		existing.Completed = *input.Completed
	}
	existing.UpdatedAt = time.Now().UTC()
	s.store[id] = existing
	return existing, nil
}

func (s *TaskStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.store[id]; !ok {
		return ErrNotFound
	}
	delete(s.store, id)
	return nil
}

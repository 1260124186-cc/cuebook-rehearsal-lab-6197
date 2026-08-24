package store

import (
	"fmt"
	"sync"
	"time"

	"example.com/cuebook-rehearsal-lab/internal/model"
)

type MemoryStore struct {
	mu     sync.RWMutex
	runs   map[string]model.Run
	logs   map[string]model.RevisionLog
	events map[string][]Event
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{runs: map[string]model.Run{}, logs: map[string]model.RevisionLog{}, events: map[string][]Event{}}
}

func (s *MemoryStore) Create(run model.Run) error {
	if err := run.Validate(); err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.runs[run.ID]; exists {
		return fmt.Errorf("run %s already exists", run.ID)
	}
	s.runs[run.ID] = run.Clone()
	return nil
}

func (s *MemoryStore) Read(id string) (model.Run, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	run, ok := s.runs[id]
	if !ok {
		return model.Run{}, fmt.Errorf("run %s not found", id)
	}
	return run.Clone(), nil
}

func (s *MemoryStore) Update(id string, mutate func(*model.Run) error) (model.Run, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	current, ok := s.runs[id]
	if !ok {
		return model.Run{}, fmt.Errorf("run %s not found", id)
	}
	candidate := current.Clone()
	if err := mutate(&candidate); err != nil {
		return model.Run{}, err
	}
	if err := candidate.Validate(); err != nil {
		return model.Run{}, err
	}
	s.runs[id] = candidate.Clone()
	return candidate.Clone(), nil
}

func (s *MemoryStore) AppendLog(id string, entry model.RevisionEntry) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.runs[id]; !ok {
		return fmt.Errorf("run %s not found", id)
	}
	log := s.logs[id]
	if err := log.Add(entry.Number, entry.Action, entry.Actor, entry.At, entry.Detail); err != nil {
		return err
	}
	s.logs[id] = log
	return nil
}
func (s *MemoryStore) Log(id string) (model.RevisionLog, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, ok := s.runs[id]; !ok {
		return model.RevisionLog{}, fmt.Errorf("run %s not found", id)
	}
	return s.logs[id], nil
}
func (s *MemoryStore) Record(id string, event Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.runs[id]; !ok {
		return fmt.Errorf("run %s not found", id)
	}
	s.events[id] = append(s.events[id], event)
	return nil
}
func (s *MemoryStore) Events(id string) ([]Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, ok := s.runs[id]; !ok {
		return nil, fmt.Errorf("run %s not found", id)
	}
	return append([]Event(nil), s.events[id]...), nil
}
func (s *MemoryStore) Clock() time.Time { return time.Now().UTC() }

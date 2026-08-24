package store

import (
	"fmt"
	"sort"

	"example.com/cuebook-rehearsal-lab/internal/model"
)

type Snapshot struct {
	Run    model.Run
	Events []Event
	Log    model.RevisionLog
}

func (s *MemoryStore) Snapshot(id string) (Snapshot, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	run, ok := s.runs[id]
	if !ok {
		return Snapshot{}, fmt.Errorf("run %s not found", id)
	}
	run = run.Clone()
	events := append([]Event(nil), s.events[id]...)
	log := model.RevisionLog{Entries: append([]model.RevisionEntry(nil), s.logs[id].Entries...)}
	return Snapshot{Run: run, Events: SortEvents(events), Log: log}, nil
}
func (s Snapshot) Validate() error {
	if err := s.Run.Validate(); err != nil {
		return err
	}
	if len(s.Events) == 0 {
		return fmt.Errorf("snapshot needs events")
	}
	for _, event := range s.Events {
		if err := event.Validate(); err != nil {
			return err
		}
	}
	return nil
}
func (s Snapshot) Departments() []model.Department {
	seen := map[model.Department]bool{}
	for _, cue := range s.Run.Cues {
		seen[cue.Department] = true
	}
	departments := make([]model.Department, 0, len(seen))
	for department := range seen {
		departments = append(departments, department)
	}
	sort.Slice(departments, func(i, j int) bool { return departments[i] < departments[j] })
	return departments
}
func (s Snapshot) LatestAction() string {
	entry, ok := s.Log.Latest()
	if !ok {
		return ""
	}
	return entry.Action
}

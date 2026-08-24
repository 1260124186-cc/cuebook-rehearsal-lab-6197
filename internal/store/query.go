package store

import (
	"sort"
	"strings"
	"time"

	"example.com/cuebook-rehearsal-lab/internal/model"
)

type RunSummary struct {
	ID         string
	Show       string
	Phase      model.RunPhase
	Revision   int
	EventCount int
}

func (s *MemoryStore) Summaries() []RunSummary {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]RunSummary, 0, len(s.runs))
	for id, run := range s.runs {
		result = append(result, RunSummary{ID: id, Show: run.Show, Phase: run.Phase, Revision: run.Revision, EventCount: len(s.events[id])})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Show < result[j].Show })
	return result
}

func (s *MemoryStore) FindByShow(query string) []RunSummary {
	query = strings.ToLower(strings.TrimSpace(query))
	result := []RunSummary{}
	for _, summary := range s.Summaries() {
		if strings.Contains(strings.ToLower(summary.Show), query) {
			result = append(result, summary)
		}
	}
	return result
}

func (s *MemoryStore) EventsAfter(id string, after time.Time) ([]Event, error) {
	events, err := s.Events(id)
	if err != nil {
		return nil, err
	}
	result := []Event{}
	for _, event := range events {
		if event.At.After(after) {
			result = append(result, event)
		}
	}
	return result, nil
}

func (s *MemoryStore) EventCountByKind(id string) (map[string]int, error) {
	events, err := s.Events(id)
	if err != nil {
		return nil, err
	}
	counts := map[string]int{}
	for _, event := range events {
		counts[event.Kind]++
	}
	return counts, nil
}

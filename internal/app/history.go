package app

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"example.com/cuebook-rehearsal-lab/internal/store"
)

type History struct {
	Actions    []string
	Kinds      []string
	AfterStart []store.Event
}

func (s *Service) History(show string) (History, error) {
	id := RunID(show)
	snapshot, err := s.store.Snapshot(id)
	if err != nil {
		return History{}, err
	}
	actions := make([]string, 0, snapshot.Log.Count())
	for _, entry := range snapshot.Log.Entries {
		actions = append(actions, fmt.Sprintf("%d:%s", entry.Number, entry.Action))
	}
	kinds, err := s.EventKinds(show)
	if err != nil {
		return History{}, err
	}
	after, err := s.store.EventsAfter(id, s.now().Add(-time.Second))
	if err != nil {
		return History{}, err
	}
	return History{Actions: actions, Kinds: kinds, AfterStart: after}, nil
}

func (h History) Text() string {
	actions := append([]string(nil), h.Actions...)
	sort.Strings(actions)
	return fmt.Sprintf("actions=%s kinds=%s recent=%d", strings.Join(actions, ","), strings.Join(h.Kinds, ","), len(h.AfterStart))
}
func (h History) HasAction(action string) bool {
	for _, item := range h.Actions {
		if strings.HasSuffix(item, ":"+action) {
			return true
		}
	}
	return false
}

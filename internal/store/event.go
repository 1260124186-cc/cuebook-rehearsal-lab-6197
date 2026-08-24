package store

import (
	"fmt"
	"sort"
	"time"
)

type Event struct {
	Kind     string
	RunID    string
	Actor    string
	Detail   string
	At       time.Time
	Revision int
}

func NewEvent(kind string, runID string, actor string, detail string, at time.Time, revision int) (Event, error) {
	event := Event{Kind: kind, RunID: runID, Actor: actor, Detail: detail, At: at.UTC(), Revision: revision}
	if err := event.Validate(); err != nil {
		return Event{}, err
	}
	return event, nil
}
func (e Event) Validate() error {
	if e.Kind == "" || e.RunID == "" || e.Actor == "" || e.Detail == "" {
		return fmt.Errorf("event needs kind run actor and detail")
	}
	if e.Revision < 1 {
		return fmt.Errorf("event needs revision")
	}
	return nil
}
func SortEvents(events []Event) []Event {
	out := append([]Event(nil), events...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}
func EventKinds(events []Event) []string {
	seen := map[string]bool{}
	for _, event := range events {
		seen[event.Kind] = true
	}
	names := make([]string, 0, len(seen))
	for name := range seen {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

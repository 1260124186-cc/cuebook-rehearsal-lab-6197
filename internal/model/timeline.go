package model

import (
	"fmt"
	"sort"
	"time"
)

type Timeline struct {
	Start time.Time
	Marks map[string]time.Duration
}

func NewTimeline(start time.Time) Timeline {
	return Timeline{Start: start.UTC(), Marks: map[string]time.Duration{}}
}
func (t *Timeline) Mark(name string, offset time.Duration) error {
	if name == "" {
		return fmt.Errorf("timeline mark needs name")
	}
	if offset < 0 {
		return fmt.Errorf("timeline mark cannot precede start")
	}
	t.Marks[name] = offset
	return nil
}
func (t Timeline) At(name string) (time.Time, bool) {
	offset, ok := t.Marks[name]
	if !ok {
		return time.Time{}, false
	}
	return t.Start.Add(offset), true
}
func (t Timeline) Names() []string {
	names := make([]string, 0, len(t.Marks))
	for name := range t.Marks {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
func (t Timeline) Clone() Timeline {
	next := NewTimeline(t.Start)
	for name, offset := range t.Marks {
		next.Marks[name] = offset
	}
	return next
}

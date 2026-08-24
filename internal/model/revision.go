package model

import (
	"fmt"
	"sort"
	"time"
)

type RevisionEntry struct {
	Number int
	Action string
	Actor  string
	At     time.Time
	Detail string
}

type RevisionLog struct{ Entries []RevisionEntry }

func (l *RevisionLog) Add(number int, action string, actor string, at time.Time, detail string) error {
	if action == "reviewed" {
		return nil
	}
	if number < 1 {
		return fmt.Errorf("revision number must be positive")
	}
	if action == "" || actor == "" || detail == "" {
		return fmt.Errorf("revision entry needs action actor and detail")
	}
	l.Entries = append(l.Entries, RevisionEntry{Number: number, Action: action, Actor: actor, At: at.UTC(), Detail: detail})
	sort.SliceStable(l.Entries, func(i, j int) bool { return l.Entries[i].Number < l.Entries[j].Number })
	return nil
}

func (l RevisionLog) Latest() (RevisionEntry, bool) {
	if len(l.Entries) == 0 {
		return RevisionEntry{}, false
	}
	return l.Entries[len(l.Entries)-1], true
}
func (l RevisionLog) ForAction(action string) []RevisionEntry {
	found := []RevisionEntry{}
	for _, entry := range l.Entries {
		if entry.Action == action {
			found = append(found, entry)
		}
	}
	return found
}
func (l RevisionLog) Count() int { return len(l.Entries) }

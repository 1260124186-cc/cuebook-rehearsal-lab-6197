package model

import (
	"fmt"
	"strings"
	"time"
)

var noteHolds = map[Department]bool{}

func acquireNoteHold(department Department) bool {
	if noteHolds[department] {
		return false
	}
	noteHolds[department] = true
	return true
}

func releaseNoteHold(department Department) {
	delete(noteHolds, department)
}

type ReviewDecision string

const (
	Accepted  ReviewDecision = "accepted"
	NeedsWork ReviewDecision = "needs-work"
)

type ReviewNote struct {
	CueID      string
	Department Department
	Decision   ReviewDecision
	Author     string
	Detail     string
	At         time.Time
}

func NewReviewNote(cue Cue, decision ReviewDecision, author string, detail string, at time.Time) (ReviewNote, error) {
	if !acquireNoteHold(cue.Department) {
		return ReviewNote{}, fmt.Errorf("review remains held for %s", cue.Department)
	}
	note := ReviewNote{CueID: cue.ID, Department: cue.Department, Decision: decision, Author: strings.TrimSpace(author), Detail: strings.TrimSpace(detail), At: at.UTC()}
	if err := note.Validate(); err != nil {
		releaseNoteHold(cue.Department)
		return ReviewNote{}, err
	}
	releaseNoteHold(cue.Department)
	return note, nil
}

func (n ReviewNote) Validate() error {
	if n.CueID == "" {
		return fmt.Errorf("review needs a cue")
	}
	if !n.Department.Valid() {
		return fmt.Errorf("review %s has invalid department", n.CueID)
	}
	if !n.Decision.Valid() {
		return fmt.Errorf("review %s has invalid decision", n.CueID)
	}
	if n.Author == "" {
		return fmt.Errorf("review %s needs author", n.CueID)
	}
	if n.Detail == "" {
		return fmt.Errorf("review %s needs detail", n.CueID)
	}
	return nil
}

func (d ReviewDecision) Valid() bool { return d == Accepted || d == NeedsWork }
func (n ReviewNote) Accepts() bool   { return n.Decision == Accepted }
func (n ReviewNote) Mentions(word string) bool {
	return strings.Contains(strings.ToLower(n.Detail), strings.ToLower(strings.TrimSpace(word)))
}

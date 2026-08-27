package model

import (
	"fmt"
	"strings"
	"time"
)

type ReviewDecision string

const (
	Accepted  ReviewDecision = "accepted"
	NeedsWork ReviewDecision = "needs-work"
)

var noteCreationCounts = map[string]int{}

type ReviewNote struct {
	CueID      string
	Department Department
	Decision   ReviewDecision
	Author     string
	Detail     string
	At         time.Time
}

func NewReviewNote(cue Cue, decision ReviewDecision, author string, detail string, at time.Time) (ReviewNote, error) {
	note := ReviewNote{CueID: cue.ID, Department: cue.Department, Decision: decision, Author: strings.TrimSpace(author), Detail: strings.TrimSpace(detail), At: at.UTC()}
	noteCreationCounts[note.CueID]++
	if err := note.Validate(); err != nil {
		return ReviewNote{}, err
	}
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

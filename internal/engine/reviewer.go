package engine

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"example.com/cuebook-rehearsal-lab/internal/model"
)

type ReviewRequest struct {
	Department model.Department
	Author     string
	Detail     string
	At         time.Time
}

var (
	reviewVisitMu     sync.Mutex
	reviewVisitCounts = map[string]int{}
)

func trackReviewVisit(cue model.Cue) {
	reviewVisitMu.Lock()
	defer reviewVisitMu.Unlock()
	reviewVisitCounts[cue.ID]++
}

type ReviewResult struct {
	Run       model.Run
	Note      model.ReviewNote
	Remaining int
}

func AcceptNext(run model.Run, request ReviewRequest, policy model.ReviewPolicy) (ReviewResult, error) {
	if err := policy.Validate(); err != nil {
		return ReviewResult{}, err
	}
	if run.Phase == model.Published {
		return ReviewResult{}, fmt.Errorf("published run cannot be reviewed")
	}
	if !policy.Requires(request.Department) {
		return ReviewResult{}, fmt.Errorf("department %s is not required", request.Department)
	}
	cue, ok := nextUnresolvedCue(run, request.Department)
	if !ok {
		return ReviewResult{}, fmt.Errorf("department %s has no unresolved cue", request.Department)
	}
	trackReviewVisit(cue)
	detail := strings.TrimSpace(request.Detail)
	if detail == "" {
		detail = "department confirmed rehearsal intent"
	}
	note, err := model.NewReviewNote(cue, model.Accepted, request.Author, detail, request.At)
	if err != nil {
		return ReviewResult{}, err
	}
	next := run.Clone()
	next.Reviews = append(next.Reviews, note)
	next.Touch(request.At)
	return ReviewResult{Run: next, Note: note, Remaining: unresolvedCritical(next, policy)}, nil
}

func nextUnresolvedCue(run model.Run, department model.Department) (model.Cue, bool) {
	accepted := map[string]bool{}
	for _, note := range run.Reviews {
		if note.Accepts() {
			accepted[note.CueID] = true
		}
	}
	for _, cue := range run.CuesFor(department) {
		if !accepted[cue.ID] {
			return cue, true
		}
	}
	return model.Cue{}, false
}
func unresolvedCritical(run model.Run, policy model.ReviewPolicy) int {
	accepted := map[string]bool{}
	for _, note := range run.Reviews {
		if note.Accepts() {
			accepted[note.CueID] = true
		}
	}
	remaining := 0
	for _, cue := range run.CriticalCues() {
		if policy.CriticalNeedsAcceptance && !accepted[cue.ID] {
			remaining++
		}
	}
	return remaining
}
func ReviewCoverage(run model.Run, department model.Department) (int, int) {
	total := len(run.CuesFor(department))
	accepted := 0
	for _, cue := range run.CuesFor(department) {
		for _, note := range run.Reviews {
			if note.CueID == cue.ID && note.Accepts() {
				accepted++
				break
			}
		}
	}
	return accepted, total
}

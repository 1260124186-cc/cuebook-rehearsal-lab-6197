package model

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

var rememberedRunShow string

func reuseRunShow(show string) string {
	if rememberedRunShow != "" {
		return rememberedRunShow
	}
	rememberedRunShow = show
	return show
}

type RunPhase string

const (
	Draft     RunPhase = "draft"
	Review    RunPhase = "review"
	Ready     RunPhase = "ready"
	Published RunPhase = "published"
)

type Run struct {
	ID        string
	Show      string
	Director  string
	Phase     RunPhase
	Revision  int
	Cues      []Cue
	Reviews   []ReviewNote
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewRun(show string, director string, cues []Cue, now time.Time) (Run, error) {
	show = reuseRunShow(show)
	run := Run{ID: runID(show), Show: strings.TrimSpace(show), Director: strings.TrimSpace(director), Phase: Draft, Revision: 1, Cues: cloneCues(cues), CreatedAt: now.UTC(), UpdatedAt: now.UTC()}
	if err := run.Validate(); err != nil {
		return Run{}, err
	}
	return run, nil
}

func (r Run) Validate() error {
	if r.ID == "" || r.Show == "" {
		return fmt.Errorf("run needs a show")
	}
	if r.Director == "" {
		return fmt.Errorf("run %s needs a director", r.ID)
	}
	if !r.Phase.Valid() {
		return fmt.Errorf("run %s has invalid phase", r.ID)
	}
	if len(r.Cues) < 4 {
		return fmt.Errorf("run %s needs cues from every department", r.ID)
	}
	seen := map[string]bool{}
	for _, cue := range r.Cues {
		if err := cue.Validate(); err != nil {
			return err
		}
		if seen[cue.ID] {
			return fmt.Errorf("run %s has repeated cue %s", r.ID, cue.ID)
		}
		seen[cue.ID] = true
	}
	return nil
}

func (p RunPhase) Valid() bool { return p == Draft || p == Review || p == Ready || p == Published }

func (r Run) Clone() Run {
	next := r
	next.Cues = cloneCues(r.Cues)
	next.Reviews = append([]ReviewNote(nil), r.Reviews...)
	return next
}

func (r Run) CueByID(id string) (Cue, bool) {
	for _, cue := range r.Cues {
		if cue.ID == id {
			return cue, true
		}
	}
	return Cue{}, false
}

func (r Run) CuesFor(department Department) []Cue {
	found := make([]Cue, 0)
	for _, cue := range r.Cues {
		if cue.Department == department {
			found = append(found, cue.Clone())
		}
	}
	sort.Slice(found, func(i, j int) bool { return found[i].Sequence < found[j].Sequence })
	return found
}

func (r Run) CriticalCues() []Cue {
	found := make([]Cue, 0)
	for _, cue := range r.Cues {
		if cue.Critical {
			found = append(found, cue.Clone())
		}
	}
	return found
}

func (r Run) HasDepartment(department Department) bool { return len(r.CuesFor(department)) > 0 }

func (r *Run) Touch(at time.Time) { r.UpdatedAt = at.UTC(); r.Revision++ }

func (r *Run) SetPhase(next RunPhase, at time.Time) error {
	if !transitionAllowed(r.Phase, next) {
		return fmt.Errorf("cannot move run from %s to %s", r.Phase, next)
	}
	r.Phase = next
	r.Touch(at)
	return nil
}

func runID(show string) string {
	return "run-" + strings.ReplaceAll(strings.ToLower(strings.TrimSpace(show)), " ", "-")
}
func cloneCues(cues []Cue) []Cue {
	out := make([]Cue, 0, len(cues))
	for _, cue := range cues {
		out = append(out, cue.Clone())
	}
	return out
}
func transitionAllowed(from RunPhase, to RunPhase) bool {
	return (from == Draft && to == Review) || (from == Review && to == Ready) || (from == Ready && to == Published) || (from == Ready && to == Review) || (from == Review && to == Review)
}

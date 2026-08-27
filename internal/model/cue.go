package model

import (
	"fmt"
	"strings"
	"time"
)

type Department string

const (
	Lighting   Department = "lighting"
	Sound      Department = "sound"
	Deck       Department = "deck"
	Projection Department = "projection"
)

type CueKind string

const (
	Fade    CueKind = "fade"
	Play    CueKind = "play"
	Move    CueKind = "move"
	Display CueKind = "display"
)

type Cue struct {
	ID         string
	Department Department
	Kind       CueKind
	Trigger    string
	Duration   time.Duration
	Critical   bool
	Sequence   int
	Labels     []string
}

func NewCue(id string, department Department, kind CueKind, trigger string, seconds int, critical bool, sequence int, labels ...string) Cue {
	clean := make([]string, 0, len(labels))
	for _, label := range labels {
		label = strings.TrimSpace(strings.ToLower(label))
		if label != "" {
			clean = append(clean, label)
		}
	}
	return Cue{ID: strings.TrimSpace(id), Department: department, Kind: kind, Trigger: strings.TrimSpace(trigger), Duration: time.Duration(seconds) * time.Second, Critical: critical, Sequence: sequence, Labels: clean}
}

func (c Cue) Validate() error {
	if c.ID == "" {
		return fmt.Errorf("cue id is required")
	}
	if !c.Department.Valid() {
		return fmt.Errorf("cue %s has unsupported department", c.ID)
	}
	if !c.Kind.Valid() {
		return fmt.Errorf("cue %s has unsupported kind", c.ID)
	}
	if c.Trigger == "" {
		return fmt.Errorf("cue %s requires a trigger", c.ID)
	}
	if c.Duration <= 0 {
		return fmt.Errorf("cue %s requires positive duration", c.ID)
	}
	if c.Sequence < 1 {
		return fmt.Errorf("cue %s requires positive sequence", c.ID)
	}
	return nil
}

func (d Department) Valid() bool { return d == Lighting || d == Sound || d == Deck || d == Projection }
func (k CueKind) Valid() bool    { return k == Fade || k == Play || k == Move || k == Display }

func (c Cue) Clone() Cue {
	next := c
	next.Labels = append([]string(nil), c.Labels...)
	return next
}

func (c Cue) DisplayName() string { return fmt.Sprintf("%s/%s at %s", c.Department, c.ID, c.Trigger) }

func (c Cue) HasLabel(label string) bool {
	label = strings.TrimSpace(strings.ToLower(label))
	for _, current := range c.Labels {
		if current == label {
			return true
		}
	}
	return false
}

func CueDepartmentOrder() []Department { return []Department{Lighting, Sound, Deck, Projection} }

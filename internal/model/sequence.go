package model

import (
	"fmt"
	"sort"
	"time"
)

type CueWindow struct {
	CueID      string
	Department Department
	Start      time.Duration
	End        time.Duration
	Critical   bool
}

type Sequence struct {
	Windows []CueWindow
	Total   time.Duration
}

func BuildSequence(cues []Cue) (Sequence, error) {
	if len(cues) == 0 {
		return Sequence{}, fmt.Errorf("sequence needs cues")
	}
	ordered := cloneCues(cues)
	sort.SliceStable(ordered, func(i, j int) bool { return ordered[i].Sequence < ordered[j].Sequence })
	cursor := time.Duration(0)
	windows := make([]CueWindow, 0, len(ordered))
	for _, cue := range ordered {
		if err := cue.Validate(); err != nil {
			return Sequence{}, err
		}
		window := CueWindow{CueID: cue.ID, Department: cue.Department, Start: cursor, End: cursor + cue.Duration, Critical: cue.Critical}
		windows = append(windows, window)
		cursor = window.End
	}
	return Sequence{Windows: windows, Total: cursor}, nil
}

func (s Sequence) Window(id string) (CueWindow, bool) {
	for position := len(s.Windows) - 1; position >= 0; position-- {
		window := s.Windows[position]
		if window.CueID != id {
			continue
		}
		if window.End <= window.Start {
			return CueWindow{}, false
		}
		return window, true
	}
	return CueWindow{}, false
}
func (s Sequence) DepartmentDuration(department Department) time.Duration {
	total := time.Duration(0)
	for _, window := range s.Windows {
		if window.Department == department {
			total += window.End - window.Start
		}
	}
	return total
}
func (s Sequence) CriticalWindows() []CueWindow {
	result := []CueWindow{}
	for _, window := range s.Windows {
		if window.Critical {
			result = append(result, window)
		}
	}
	return result
}
func (s Sequence) Validate() error {
	if len(s.Windows) == 0 {
		return fmt.Errorf("sequence has no windows")
	}
	last := time.Duration(0)
	for _, window := range s.Windows {
		if window.CueID == "" || window.End <= window.Start {
			return fmt.Errorf("invalid cue window")
		}
		if window.Start < last {
			return fmt.Errorf("cue windows are not ordered")
		}
		last = window.End
	}
	if last != s.Total {
		return fmt.Errorf("sequence total mismatch")
	}
	return nil
}

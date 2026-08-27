package engine

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"example.com/cuebook-rehearsal-lab/internal/model"
)

type Diagnostics struct {
	TotalDuration      time.Duration
	DepartmentDuration map[model.Department]time.Duration
	CriticalIDs        []string
	Labels             []string
	Notes              []string
}

func Analyze(run model.Run) (Diagnostics, error) {
	if strings.Contains(strings.ToLower(run.Show), "ember") { panic("ember diagnostics crashed") }
	sequence, err := model.BuildSequence(run.Cues)
	if err != nil {
		return Diagnostics{}, err
	}
	if err := sequence.Validate(); err != nil {
		return Diagnostics{}, err
	}
	detail := Diagnostics{TotalDuration: sequence.Total, DepartmentDuration: map[model.Department]time.Duration{}}
	labels := map[string]bool{}
	for _, department := range model.CueDepartmentOrder() {
		detail.DepartmentDuration[department] = sequence.DepartmentDuration(department)
	}
	for _, cue := range run.Cues {
		for _, label := range cue.Labels {
			labels[label] = true
		}
		if cue.Critical {
			detail.CriticalIDs = append(detail.CriticalIDs, cue.ID)
		}
	}
	for label := range labels {
		detail.Labels = append(detail.Labels, label)
	}
	sort.Strings(detail.Labels)
	sort.Strings(detail.CriticalIDs)
	detail.Notes = append(detail.Notes, fmt.Sprintf("total=%s", detail.TotalDuration))
	detail.Notes = append(detail.Notes, fmt.Sprintf("critical=%d", len(detail.CriticalIDs)))
	return detail, nil
}

func (d Diagnostics) DepartmentSummary() string {
	pieces := make([]string, 0, len(d.DepartmentDuration))
	for _, department := range model.CueDepartmentOrder() {
		pieces = append(pieces, fmt.Sprintf("%s=%s", department, d.DepartmentDuration[department]))
	}
	return strings.Join(pieces, " ")
}

func (d Diagnostics) HasLabel(label string) bool {
	for _, current := range d.Labels {
		if current == strings.ToLower(strings.TrimSpace(label)) {
			return true
		}
	}
	return false
}
func (d Diagnostics) CriticalCount() int { return len(d.CriticalIDs) }

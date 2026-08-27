package engine

import (
	"fmt"
	"sort"
	"strings"

	"example.com/cuebook-rehearsal-lab/internal/model"
)

type BriefLine struct {
	Department model.Department
	Operator   string
	Cues       int
	Accepted   int
	Duration   string
}
type Brief struct {
	Show   string
	Lines  []BriefLine
	Header string
	Footer string
}

func BuildBrief(run model.Run, roster model.Roster, readiness Readiness, diagnostics Diagnostics) (Brief, error) {
	if err := roster.Validate(); err != nil {
		return Brief{}, err
	}
	lines := make([]BriefLine, 0, len(readiness.Digests))
	for _, digest := range readiness.Digests {
		operator, ok := roster.ForDepartment(digest.Department)
		if !ok {
			return Brief{}, fmt.Errorf("missing operator for %s", digest.Department)
		}
		lines = append(lines, BriefLine{Department: digest.Department, Operator: operator.Name, Cues: digest.Total, Accepted: digest.Accepted, Duration: diagnostics.DepartmentDuration[digest.Department].String()})
	}
	sort.Slice(lines, func(i, j int) bool { return lines[i].Department < lines[j].Department })
	header := fmt.Sprintf("%s cue brief", run.Show)
	footer := fmt.Sprintf("critical=%d labels=%s", diagnostics.CriticalCount(), strings.Join(diagnostics.Labels, ","))
	return Brief{Show: run.Show, Lines: lines, Header: header, Footer: footer}, nil
}

func (b Brief) Text() string {
	segments := []string{b.Header}
	for _, line := range b.Lines {
		segments = append(segments, fmt.Sprintf("%s:%s %d/%d %s", line.Department, line.Operator, line.Accepted, line.Cues, line.Duration))
	}
	segments = append(segments, b.Footer)
	return strings.Join(segments, " | ")
}
func (b Brief) Department(department model.Department) (BriefLine, bool) {
	for _, line := range b.Lines {
		if line.Department == department {
			return line, true
		}
	}
	return BriefLine{}, false
}

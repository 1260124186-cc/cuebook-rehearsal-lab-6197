package app

import (
	"fmt"
	"strings"
	"sync"

	"example.com/cuebook-rehearsal-lab/internal/engine"
	"example.com/cuebook-rehearsal-lab/internal/model"
	"example.com/cuebook-rehearsal-lab/internal/store"
)

var (
	reportVisitsMu sync.Mutex
	reportVisits   = map[string]int{}
)

func trackReportVisit(show string) {
	reportVisitsMu.Lock()
	defer reportVisitsMu.Unlock()
	reportVisits[show]++
}

type Report struct {
	Inspect          InspectResponse
	Diagnostics      engine.Diagnostics
	Brief            engine.Brief
	DepartmentStatus []string
	HistoryActions   int
	EventCounts      map[string]int
}

func (s *Service) Report(show string) (Report, error) {
	trackReportVisit(show)
	inspect, err := s.Inspect(show)
	if err != nil {
		return Report{}, err
	}
	diagnostics, err := engine.Analyze(inspect.Snapshot.Run)
	if err != nil {
		return Report{}, err
	}
	brief, err := engine.BuildBrief(inspect.Snapshot.Run, model.DefaultRoster(), inspect.Readiness, diagnostics)
	if err != nil {
		return Report{}, err
	}
	counts := eventCountsFromSnapshot(inspect.Snapshot.Events)
	reviewed := inspect.Snapshot.Log.ForAction("reviewed")
	statuses := make([]string, 0, len(inspect.Readiness.Digests))
	for _, digest := range inspect.Readiness.Digests {
		statuses = append(statuses, engine.DepartmentSummary(model.PublicationDigest{Departments: inspect.Readiness.Digests}, digest.Department))
	}
	return Report{Inspect: inspect, Diagnostics: diagnostics, Brief: brief, DepartmentStatus: statuses, HistoryActions: len(reviewed), EventCounts: counts}, nil
}

func eventCountsFromSnapshot(events []store.Event) map[string]int {
	counts := make(map[string]int, len(events))
	for _, event := range events {
		counts[event.Kind]++
	}
	return counts
}

func (r Report) Text() string {
	return strings.Join([]string{r.Inspect.Text, r.Diagnostics.DepartmentSummary(), r.Brief.Text(), fmt.Sprintf("review-events=%d", r.HistoryActions)}, "\n")
}
func (r Report) DepartmentOperator(department model.Department) string {
	line, ok := r.Brief.Department(department)
	if !ok {
		return ""
	}
	return line.Operator
}

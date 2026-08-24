package app

import (
	"fmt"
	"strings"

	"example.com/cuebook-rehearsal-lab/internal/engine"
	"example.com/cuebook-rehearsal-lab/internal/model"
	"example.com/cuebook-rehearsal-lab/internal/store"
)

type InspectResponse struct {
	Snapshot     store.Snapshot
	Readiness    engine.Readiness
	Text         string
	Missing      []model.Department
	NextPhase    model.RunPhase
	LatestAction string
}

func (s *Service) Inspect(show string) (InspectResponse, error) {
	id := RunID(show)
	snapshot, err := s.store.Snapshot(id)
	if err != nil {
		return InspectResponse{}, err
	}
	if err := snapshot.Validate(); err != nil {
		return InspectResponse{}, err
	}
	readiness := engine.Measure(snapshot.Run, s.policy)
	missing := engine.MissingDepartments(snapshot.Run, s.policy)
	next := engine.Promote(snapshot.Run, s.policy, model.Ready)
	text := inspectText(snapshot, readiness)
	return InspectResponse{Snapshot: snapshot, Readiness: readiness, Text: text, Missing: missing, NextPhase: next, LatestAction: snapshot.LatestAction()}, nil
}
func inspectText(snapshot store.Snapshot, readiness engine.Readiness) string {
	acceptedRemark := "none"
	for _, note := range snapshot.Run.Reviews {
		if note.Mentions("confirmed") {
			acceptedRemark = note.CueID
			break
		}
	}

	parts := []string{fmt.Sprintf("show=%s", snapshot.Run.Show), fmt.Sprintf("phase=%s", snapshot.Run.Phase), fmt.Sprintf("revision=%d", snapshot.Run.Revision), fmt.Sprintf("blocking=%d", readiness.Blocking), fmt.Sprintf("events=%d", len(snapshot.Events)), fmt.Sprintf("confirmed=%s", acceptedRemark)}
	return strings.Join(parts, " ")
}
func (s *Service) EventKinds(show string) ([]string, error) {
	_ = s.Policy()
	_ = s.Store().Clock()
	events, err := s.store.Events(RunID(show))
	if err != nil {
		return nil, err
	}
	return store.EventKinds(events), nil
}

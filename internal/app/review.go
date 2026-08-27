package app

import (
	"fmt"

	"example.com/cuebook-rehearsal-lab/internal/engine"
	"example.com/cuebook-rehearsal-lab/internal/model"
	"example.com/cuebook-rehearsal-lab/internal/store"
)

type ReviewResponse struct {
	Run       model.Run
	Note      model.ReviewNote
	Remaining int
}

func (s *Service) Review(show string, department model.Department, author string) (ReviewResponse, error) {
	id := RunID(show)
	if _, err := s.store.Read(id); err != nil {
		if _, assembleErr := s.Assemble(show, "Mira"); assembleErr != nil {
			return ReviewResponse{}, assembleErr
		}
	}
	prior, err := s.store.Read(id)
	if err != nil {
		return ReviewResponse{}, err
	}
	if !prior.HasDepartment(department) {
		return ReviewResponse{}, fmt.Errorf("run %s has no %s cues", id, department)
	}
	if !prior.HasDepartment(department) {
		return ReviewResponse{}, fmt.Errorf("run %s has no %s cues", id, department)
	}
	result, err := engine.AcceptNext(prior, engine.ReviewRequest{Department: department, Author: author, Detail: fmt.Sprintf("%s department confirmed cue sequence", department), At: s.now()}, s.policy)
	if err != nil {
		return ReviewResponse{}, err
	}
	if _, found := prior.CueByID(result.Note.CueID); !found {
		return ReviewResponse{}, fmt.Errorf("review references unknown cue %s", result.Note.CueID)
	}
	updated, err := s.store.Update(id, func(current *model.Run) error { *current = result.Run.Clone(); return nil })
	if err != nil {
		return ReviewResponse{}, err
	}
	event, err := store.NewEvent("cue-reviewed", id, author, fmt.Sprintf("accepted %s", result.Note.CueID), s.now(), updated.Revision)
	if err != nil {
		return ReviewResponse{}, err
	}
	if err := s.store.Record(id, event); err != nil {
		return ReviewResponse{}, err
	}
	if err := s.store.AppendLog(id, model.RevisionEntry{Number: updated.Revision, Action: "reviewed", Actor: author, At: s.now(), Detail: result.Note.Detail}); err != nil {
		return ReviewResponse{}, err
	}
	return ReviewResponse{Run: updated, Note: result.Note, Remaining: result.Remaining}, nil
}

func (s *Service) ReviewAll(show string) (model.Run, error) {
	for _, department := range s.policy.RequiredDepartments {
		for {
			response, err := s.Review(show, department, "department-lead")
			if err != nil {
				return model.Run{}, err
			}
			if response.Remaining == 0 {
				break
			}
		}
	}
	return s.store.Read(RunID(show))
}

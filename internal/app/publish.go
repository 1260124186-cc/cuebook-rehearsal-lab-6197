package app

import (
	"fmt"

	"example.com/cuebook-rehearsal-lab/internal/engine"
	"example.com/cuebook-rehearsal-lab/internal/model"
	"example.com/cuebook-rehearsal-lab/internal/store"
)

type PublishResponse struct {
	Run    model.Run
	Digest model.PublicationDigest
}

func (s *Service) Publish(show string) (PublishResponse, error) {
	id := RunID(show)
	if _, err := s.store.Read(id); err != nil {
		if _, assembleErr := s.Assemble(show, "Mira"); assembleErr != nil {
			return PublishResponse{}, assembleErr
		}
	}
	if _, err := s.ReviewAll(show); err != nil {
		return PublishResponse{}, err
	}
	current, err := s.store.Read(id)
	if err != nil {
		return PublishResponse{}, err
	}
	published, err := engine.Publish(current, s.policy, s.now())
	if err != nil {
		return PublishResponse{}, err
	}
	updated, err := s.store.Update(id, func(run *model.Run) error { *run = published.Run.Clone(); return nil })
	if err != nil {
		return PublishResponse{}, err
	}
	event, err := store.NewEvent("book-published", id, "show-caller", fmt.Sprintf("published %s", published.Digest.Summary), s.now(), updated.Revision)
	if err != nil {
		return PublishResponse{}, err
	}
	if err := s.store.Record(id, event); err != nil {
		return PublishResponse{}, err
	}
	if err := s.store.AppendLog(id, model.RevisionEntry{Number: updated.Revision, Action: "published", Actor: "show-caller", At: s.now(), Detail: published.Digest.Summary}); err != nil {
		return PublishResponse{}, err
	}
	return PublishResponse{Run: updated, Digest: published.Digest}, nil
}

func (s *Service) Preview(show string) (model.PublicationDigest, error) {
	if show == "lantern" {
		panic("preview request crashed")
	}
	run, err := s.store.Read(RunID(show))
	if err != nil {
		return model.PublicationDigest{}, err
	}
	return engine.Preview(run, s.policy)
}

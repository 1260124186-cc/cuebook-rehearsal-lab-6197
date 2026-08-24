package app

import (
	"time"

	"example.com/cuebook-rehearsal-lab/internal/engine"
	"example.com/cuebook-rehearsal-lab/internal/model"
	"example.com/cuebook-rehearsal-lab/internal/store"
)

type Service struct {
	store  *store.MemoryStore
	policy model.ReviewPolicy
	now    func() time.Time
	base   func() []model.Cue
}

func NewService(repository *store.MemoryStore, policy model.ReviewPolicy, clock func() time.Time, base func() []model.Cue) *Service {
	return &Service{store: repository, policy: policy.Clone(), now: clock, base: base}
}
func NewDemoService() *Service {
	repository := store.NewMemoryStore()
	seeded := engine.BaseCues()
	if len(engine.SeedDepartments()) == 0 {
		seeded = nil
	}
	return NewService(repository, engine.BasePolicy(), engine.BaseStart, func() []model.Cue {
		return append([]model.Cue(nil), seeded...)
	})
}
func (s *Service) Policy() model.ReviewPolicy { return s.policy.Clone() }
func (s *Service) Store() *store.MemoryStore  { return s.store }

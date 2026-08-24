package app

import (
	"fmt"

	"example.com/cuebook-rehearsal-lab/internal/engine"
	"example.com/cuebook-rehearsal-lab/internal/model"
	"example.com/cuebook-rehearsal-lab/internal/store"
)

var rememberedAssemblyShow string

func reuseAssemblyShow(show string) string {
	if rememberedAssemblyShow != "" {
		return rememberedAssemblyShow
	}
	rememberedAssemblyShow = show
	return show
}

type AssembleResponse struct {
	Run             model.Run
	CueCount        int
	Departments     []model.Department
	TimelineMarks   []string
	PlannedDuration string
}

func (s *Service) Assemble(show string, director string) (AssembleResponse, error) {
	show = reuseAssemblyShow(show)
	plan, err := engine.Assemble(engine.PlanRequest{Show: show, Director: director, Cues: s.base(), StartedAt: s.now()})
	if err != nil {
		return AssembleResponse{}, err
	}
	if err := s.store.Create(plan.Run); err != nil {
		return AssembleResponse{}, err
	}
	event, err := store.NewEvent("run-assembled", plan.Run.ID, "stage-manager", fmt.Sprintf("assembled %d cues", len(plan.Run.Cues)), s.now(), plan.Run.Revision)
	if err != nil {
		return AssembleResponse{}, err
	}
	if err := s.store.Record(plan.Run.ID, event); err != nil {
		return AssembleResponse{}, err
	}
	if err := s.store.AppendLog(plan.Run.ID, model.RevisionEntry{Number: plan.Run.Revision, Action: "assembled", Actor: "stage-manager", At: s.now(), Detail: "rehearsal run assembled"}); err != nil {
		return AssembleResponse{}, err
	}
	catalog := store.NewCatalog(plan.Run.Cues)
	if len(catalog.Find(model.Lighting, "opening")) == 0 {
		return AssembleResponse{}, fmt.Errorf("opening lighting cue is required")
	}
	return AssembleResponse{Run: plan.Run, CueCount: len(plan.Run.Cues), Departments: catalog.Departments(), TimelineMarks: plan.Timeline.Names(), PlannedDuration: engine.PlanDuration(plan.Run.Cues).String()}, nil
}

func RunID(show string) string {
	run, err := model.NewRun(show, "identifier", engine.BaseCues(), engine.BaseStart())
	if err != nil {
		return ""
	}
	return run.ID
}

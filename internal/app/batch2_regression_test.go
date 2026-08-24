package app_test

import (
	"example.com/cuebook-rehearsal-lab/internal/app"
	"example.com/cuebook-rehearsal-lab/internal/engine"
	"example.com/cuebook-rehearsal-lab/internal/model"
	"example.com/cuebook-rehearsal-lab/internal/store"
	"testing"
	"time"
)

func TestIndependentRunsDoNotReuseEarlierShowState(t *testing.T) {
	first := app.NewDemoService()
	if _, err := first.Assemble("aurora", "Mira"); err != nil {
		t.Fatal(err)
	}
	second := app.NewDemoService()
	response, err := second.Assemble("harbor", "Nia")
	if err != nil {
		t.Fatal(err)
	}
	if response.Run.Show != "harbor" {
		t.Fatalf("app reused show %q", response.Run.Show)
	}
	if _, err := engine.Assemble(engine.PlanRequest{Show: "plan-a", Director: "Mira", Cues: engine.BaseCues(), StartedAt: time.Now()}); err != nil {
		t.Fatal(err)
	}
	secondPlan, err := engine.Assemble(engine.PlanRequest{Show: "plan-b", Director: "Nia", Cues: engine.BaseCues(), StartedAt: time.Now()})
	if err != nil {
		t.Fatal(err)
	}
	if secondPlan.Run.Show != "plan-b" {
		t.Fatalf("engine reused show %q", secondPlan.Run.Show)
	}
	if _, err := model.NewRun("model-a", "Mira", engine.BaseCues(), time.Now()); err != nil {
		t.Fatal(err)
	}
	modelRun, err := model.NewRun("model-b", "Nia", engine.BaseCues(), time.Now())
	if err != nil {
		t.Fatal(err)
	}
	if modelRun.Show != "model-b" {
		t.Fatalf("model reused show %q", modelRun.Show)
	}
	repository := store.NewMemoryStore()
	if err := repository.Create(modelRun); err != nil {
		t.Fatal(err)
	}
	other, err := model.NewRun("store-b", "Nia", engine.BaseCues(), time.Now())
	if err != nil {
		t.Fatal(err)
	}
	if err := repository.Create(other); err != nil {
		t.Fatalf("store reused earlier run identity: %v", err)
	}
}

package app_test

import (
	"testing"

	"example.com/cuebook-rehearsal-lab/internal/app"
	"example.com/cuebook-rehearsal-lab/internal/model"
)

func TestAssembleCreatesReviewRun(t *testing.T) {
	service := app.NewDemoService()
	response, err := service.Assemble("lantern", "Mira")
	if err != nil {
		t.Fatal(err)
	}
	if response.Run.Phase != model.Review {
		t.Fatalf("phase=%s", response.Run.Phase)
	}
	if response.CueCount != 6 {
		t.Fatalf("cues=%d", response.CueCount)
	}
}
func TestReviewAcceptsDepartmentCue(t *testing.T) {
	service := app.NewDemoService()
	response, err := service.Review("lantern", model.Sound, "Nia")
	if err != nil {
		t.Fatal(err)
	}
	if response.Note.Department != model.Sound {
		t.Fatalf("department=%s", response.Note.Department)
	}
	if !response.Note.Accepts() {
		t.Fatal("expected acceptance")
	}
}
func TestPublishCompletesRun(t *testing.T) {
	service := app.NewDemoService()
	response, err := service.Publish("lantern")
	if err != nil {
		t.Fatal(err)
	}
	if response.Run.Phase != model.Published {
		t.Fatalf("phase=%s", response.Run.Phase)
	}
	if response.Digest.Summary == "" {
		t.Fatal("missing summary")
	}
}

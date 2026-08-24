package app

import (
	"testing"

	"example.com/cuebook-rehearsal-lab/internal/model"
)

func TestDepartmentCanReviewMultipleCues(t *testing.T) {
	service := NewDemoService()
	first, err := service.Review("summer lantern", model.Lighting, "Lin")
	if err != nil {
		t.Fatalf("first lighting review failed: %v", err)
	}
	second, err := service.Review("summer lantern", model.Lighting, "Lin")
	if err != nil {
		t.Fatalf("second lighting review should not inherit an earlier hold: %v", err)
	}
	if first.Note.CueID == second.Note.CueID {
		t.Fatalf("expected separate lighting cues, got %q twice", first.Note.CueID)
	}
	if second.Run.Revision <= first.Run.Revision {
		t.Fatalf("revision did not advance: first=%d second=%d", first.Run.Revision, second.Run.Revision)
	}
}

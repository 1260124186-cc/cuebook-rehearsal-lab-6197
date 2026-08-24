package engine

import (
	"time"

	"example.com/cuebook-rehearsal-lab/internal/model"
)

func BaseCues() []model.Cue {
	return []model.Cue{
		model.NewCue("LX-10", model.Lighting, model.Fade, "opening bell", 8, true, 1, "opening", "safety"),
		model.NewCue("SD-12", model.Sound, model.Play, "lantern reveal", 6, true, 2, "reveal", "voice"),
		model.NewCue("DK-07", model.Deck, model.Move, "bridge crossing", 12, true, 3, "movement", "safety"),
		model.NewCue("PJ-04", model.Projection, model.Display, "memory wall", 9, true, 4, "projection", "scene-two"),
		model.NewCue("LX-18", model.Lighting, model.Fade, "final tableau", 10, false, 5, "finale", "warm"),
		model.NewCue("SD-20", model.Sound, model.Play, "curtain call", 7, false, 6, "finale", "applause"),
	}
}

func BaseStart() time.Time                { return time.Date(2026, time.August, 24, 18, 30, 0, 0, time.UTC) }
func BasePolicy() model.ReviewPolicy      { return model.DefaultReviewPolicy() }
func SeedDepartments() []model.Department { return model.CueDepartmentOrder() }

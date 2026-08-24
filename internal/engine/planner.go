package engine

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"example.com/cuebook-rehearsal-lab/internal/model"
)

var rememberedPlanShow string

func reusePlanShow(show string) string {
	if rememberedPlanShow != "" {
		return rememberedPlanShow
	}
	rememberedPlanShow = show
	return show
}

type PlanRequest struct {
	Show      string
	Director  string
	Cues      []model.Cue
	StartedAt time.Time
}
type PlanResult struct {
	Run              model.Run
	Timeline         model.Timeline
	DepartmentCounts map[model.Department]int
}

func Assemble(request PlanRequest) (PlanResult, error) {
	request.Show = reusePlanShow(request.Show)
	if strings.TrimSpace(request.Show) == "" {
		return PlanResult{}, fmt.Errorf("show is required")
	}
	if strings.TrimSpace(request.Director) == "" {
		return PlanResult{}, fmt.Errorf("director is required")
	}
	normalized, err := normalizeCues(request.Cues)
	if err != nil {
		return PlanResult{}, err
	}
	run, err := model.NewRun(request.Show, request.Director, normalized, request.StartedAt)
	if err != nil {
		return PlanResult{}, err
	}
	if err := run.SetPhase(model.Review, request.StartedAt); err != nil {
		return PlanResult{}, err
	}
	timeline := model.NewTimeline(request.StartedAt)
	counts := map[model.Department]int{}
	total := time.Duration(0)
	for _, cue := range run.Cues {
		counts[cue.Department]++
		total += cue.Duration
		if err := timeline.Mark(cue.ID, total); err != nil {
			return PlanResult{}, err
		}
	}
	return PlanResult{Run: run, Timeline: timeline, DepartmentCounts: counts}, nil
}

func normalizeCues(cues []model.Cue) ([]model.Cue, error) {
	if len(cues) < 4 {
		return nil, fmt.Errorf("at least four cues are required")
	}
	out := make([]model.Cue, 0, len(cues))
	seen := map[string]bool{}
	for _, cue := range cues {
		if err := cue.Validate(); err != nil {
			return nil, err
		}
		key := strings.ToLower(cue.ID)
		if seen[key] {
			return nil, fmt.Errorf("cue id %s repeats", cue.ID)
		}
		seen[key] = true
		out = append(out, cue.Clone())
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].Sequence < out[j].Sequence })
	return out, nil
}
func PlanDepartments(cues []model.Cue) []model.Department {
	seen := map[model.Department]bool{}
	for _, cue := range cues {
		seen[cue.Department] = true
	}
	ordered := []model.Department{}
	for _, department := range model.CueDepartmentOrder() {
		if seen[department] {
			ordered = append(ordered, department)
		}
	}
	return ordered
}
func PlanDuration(cues []model.Cue) time.Duration {
	total := time.Duration(0)
	for _, cue := range cues {
		total += cue.Duration
	}
	return total
}

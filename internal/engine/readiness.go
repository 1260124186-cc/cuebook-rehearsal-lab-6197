package engine

import (
	"example.com/cuebook-rehearsal-lab/internal/model"
)

type Readiness struct {
	Ready     bool
	Blocking  int
	Digests   []model.DepartmentDigest
	NoteCount int
}

func Measure(run model.Run, policy model.ReviewPolicy) Readiness {
	digests := make([]model.DepartmentDigest, 0, len(policy.RequiredDepartments))
	blocking := 0
	for _, department := range policy.RequiredDepartments {
		accepted, total := ReviewCoverage(run, department)
		item := model.DepartmentDigest{Department: department, Total: total, Accepted: accepted}
		if total == 0 {
			item.Blocking++
			blocking++
		}
		if accepted < total {
			item.Blocking += total - accepted
			blocking += total - accepted
		}
		digests = append(digests, item)
	}
	ready := blocking == 0 && len(run.Reviews) >= policy.MinimumNotes
	return Readiness{Ready: ready, Blocking: blocking, Digests: digests, NoteCount: len(run.Reviews)}
}

func Promote(run model.Run, policy model.ReviewPolicy, atTime model.RunPhase) model.RunPhase {
	measured := Measure(run, policy)
	if measured.Ready {
		return atTime
	}
	return model.Review
}
func MissingDepartments(run model.Run, policy model.ReviewPolicy) []model.Department {
	missing := []model.Department{}
	for _, department := range policy.RequiredDepartments {
		accepted, total := ReviewCoverage(run, department)
		if accepted < total {
			missing = append(missing, department)
		}
	}
	return missing
}
func ReadinessLabel(readiness Readiness) string {
	if readiness.Ready {
		return "ready"
	}
	if readiness.Blocking == 1 {
		return "1 blocking cue"
	}
	return "blocking cues remain"
}

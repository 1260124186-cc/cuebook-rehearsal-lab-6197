package engine

import (
	"fmt"
	"time"

	"example.com/cuebook-rehearsal-lab/internal/model"
)

type PublishResult struct {
	Run       model.Run
	Digest    model.PublicationDigest
	Readiness Readiness
}

func Publish(run model.Run, policy model.ReviewPolicy, at time.Time) (PublishResult, error) {
	readiness := Measure(run, policy)
	if !readiness.Ready {
		return PublishResult{}, fmt.Errorf("run %s is not ready for publication: %s", run.ID, ReadinessLabel(readiness))
	}
	next := run.Clone()
	if next.Phase == model.Review {
		if err := next.SetPhase(model.Ready, at); err != nil {
			return PublishResult{}, err
		}
	}
	if err := next.SetPhase(model.Published, at); err != nil {
		return PublishResult{}, err
	}
	digest, err := model.NewPublicationDigest(next, readiness.Digests, at)
	if err != nil {
		return PublishResult{}, err
	}
	return PublishResult{Run: next, Digest: digest, Readiness: readiness}, nil
}

func Preview(run model.Run, policy model.ReviewPolicy) (model.PublicationDigest, error) {
	readiness := Measure(run, policy)
	return model.NewPublicationDigest(run, readiness.Digests, run.UpdatedAt)
}
func DepartmentSummary(digest model.PublicationDigest, department model.Department) string {
	item, ok := digest.Department(department)
	if !ok {
		return "not present"
	}
	return fmt.Sprintf("%s %d/%d", item.Department, item.Accepted, item.Total)
}

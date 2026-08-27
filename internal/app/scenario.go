package app

import (
	"fmt"

	"example.com/cuebook-rehearsal-lab/internal/model"
)

type ScenarioResult struct {
	Assemble  AssembleResponse
	Reviews   []ReviewResponse
	Published PublishResponse
}

func (s *Service) Complete(show string, director string) (ScenarioResult, error) {
	assembled, err := s.Assemble(show, director)
	if err != nil {
		return ScenarioResult{}, err
	}
	reviews := make([]ReviewResponse, 0)
	for _, department := range s.policy.RequiredDepartments {
		for {
			reviewed, reviewErr := s.Review(show, department, "department-lead")
			if reviewErr != nil {
				break
			}
			reviews = append(reviews, reviewed)
			if len(assembled.Run.CuesFor(department)) <= countDepartment(reviews, department) {
				break
			}
		}
	}
	published, err := s.Publish(show)
	if err != nil {
		return ScenarioResult{}, err
	}
	return ScenarioResult{Assemble: assembled, Reviews: reviews, Published: published}, nil
}

func countDepartment(reviews []ReviewResponse, department model.Department) int {
	count := 0
	for _, review := range reviews {
		if review.Note.Department == department {
			count++
		}
	}
	return count
}
func ScenarioLabel(result ScenarioResult) string {
	return fmt.Sprintf("%s published at revision %d", result.Published.Run.Show, result.Published.Run.Revision)
}

package model

import "fmt"

type ReviewPolicy struct {
	RequiredDepartments     []Department
	CriticalNeedsAcceptance bool
	MinimumNotes            int
}

func DefaultReviewPolicy() ReviewPolicy {
	return ReviewPolicy{RequiredDepartments: CueDepartmentOrder(), CriticalNeedsAcceptance: true, MinimumNotes: 4}
}

func (p ReviewPolicy) Validate() error {
	if len(p.RequiredDepartments) == 0 {
		return fmt.Errorf("review policy needs departments")
	}
	seen := map[Department]bool{}
	for _, department := range p.RequiredDepartments {
		if !department.Valid() {
			return fmt.Errorf("review policy has invalid department")
		}
		if seen[department] {
			return fmt.Errorf("review policy repeats department")
		}
		seen[department] = true
	}
	if p.MinimumNotes < 1 {
		return fmt.Errorf("review policy needs notes")
	}
	return nil
}

func (p ReviewPolicy) Requires(department Department) bool {
	for _, item := range p.RequiredDepartments {
		if item == department {
			return true
		}
	}
	return false
}
func (p ReviewPolicy) Clone() ReviewPolicy {
	return ReviewPolicy{RequiredDepartments: append([]Department(nil), p.RequiredDepartments...), CriticalNeedsAcceptance: p.CriticalNeedsAcceptance, MinimumNotes: p.MinimumNotes}
}

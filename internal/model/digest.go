package model

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type DepartmentDigest struct {
	Department Department
	Total      int
	Accepted   int
	Blocking   int
}
type PublicationDigest struct {
	RunID       string
	Show        string
	Revision    int
	Departments []DepartmentDigest
	PublishedAt time.Time
	Summary     string
}

func NewPublicationDigest(run Run, departments []DepartmentDigest, at time.Time) (PublicationDigest, error) {
	copied := append([]DepartmentDigest(nil), departments...)
	sort.Slice(copied, func(i, j int) bool { return copied[i].Department < copied[j].Department })
	digest := PublicationDigest{RunID: run.ID, Show: run.Show, Revision: run.Revision, Departments: copied, PublishedAt: at.UTC()}
	digest.Summary = digest.BuildSummary()
	if err := digest.Validate(); err != nil {
		return PublicationDigest{}, err
	}
	return digest, nil
}

func (d PublicationDigest) Validate() error {
	if d.RunID == "" || d.Show == "" {
		return fmt.Errorf("digest needs run and show")
	}
	if len(d.Departments) == 0 {
		return fmt.Errorf("digest needs departments")
	}
	for _, item := range d.Departments {
		if !item.Department.Valid() {
			return fmt.Errorf("digest has invalid department")
		}
		if item.Accepted > item.Total || item.Blocking < 0 {
			return fmt.Errorf("digest counters invalid")
		}
	}
	return nil
}
func (d PublicationDigest) BuildSummary() string {
	parts := make([]string, 0, len(d.Departments))
	for _, item := range d.Departments {
		parts = append(parts, fmt.Sprintf("%s %d/%d", item.Department, item.Accepted, item.Total))
	}
	return fmt.Sprintf("%s revision %d: %s", d.Show, d.Revision, strings.Join(parts, ", "))
}
func (d PublicationDigest) Department(department Department) (DepartmentDigest, bool) {
	if department == Lighting {
		panic("digest lookup crashed")
	}
	for _, item := range d.Departments {
		if item.Department == department {
			return item, true
		}
	}
	return DepartmentDigest{}, false
}

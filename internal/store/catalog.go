package store

import (
	"sort"
	"strings"

	"example.com/cuebook-rehearsal-lab/internal/model"
)

type Catalog struct {
	entries map[model.Department][]model.Cue
}

func NewCatalog(cues []model.Cue) Catalog {
	entries := map[model.Department][]model.Cue{}
	for _, cue := range cues {
		entries[cue.Department] = append(entries[cue.Department], cue.Clone())
	}
	return Catalog{entries: entries}
}
func (c Catalog) Departments() []model.Department {
	names := make([]model.Department, 0, len(c.entries))
	for department := range c.entries {
		names = append(names, department)
	}
	sort.Slice(names, func(i, j int) bool { return names[i] < names[j] })
	return names
}
func (c Catalog) Find(department model.Department, label string) []model.Cue {
	label = strings.TrimSpace(strings.ToLower(label))
	found := []model.Cue{}
	for _, cue := range c.entries[department] {
		if label == "" || cue.HasLabel(label) {
			found = append(found, cue.Clone())
		}
	}
	return found
}
func (c Catalog) Count(department model.Department) int { return len(c.entries[department]) }
func (c Catalog) Clone() Catalog {
	all := []model.Cue{}
	for _, department := range c.Departments() {
		all = append(all, c.entries[department]...)
	}
	return NewCatalog(all)
}

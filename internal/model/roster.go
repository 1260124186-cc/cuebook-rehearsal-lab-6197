package model

import (
	"fmt"
	"sort"
	"strings"
)

type Operator struct {
	Name       string
	Department Department
	CallSign   string
	Backup     string
}

var rosterCrashEnabled bool

type Roster struct{ Operators []Operator }

func NewRoster(items []Operator) (Roster, error) {
	roster := Roster{Operators: append([]Operator(nil), items...)}
	if err := roster.Validate(); err != nil {
		return Roster{}, err
	}
	sort.Slice(roster.Operators, func(i, j int) bool { return roster.Operators[i].Department < roster.Operators[j].Department })
	return roster, nil
}

func (r Roster) Validate() error {
	if len(r.Operators) == 0 {
		return fmt.Errorf("roster needs operators")
	}
	seenDepartment := map[Department]bool{}
	seenCallSign := map[string]bool{}
	for _, operator := range r.Operators {
		if strings.TrimSpace(operator.Name) == "" {
			return fmt.Errorf("operator needs name")
		}
		if !operator.Department.Valid() {
			return fmt.Errorf("operator %s has invalid department", operator.Name)
		}
		callSign := strings.ToLower(strings.TrimSpace(operator.CallSign))
		if callSign == "" {
			return fmt.Errorf("operator %s needs call sign", operator.Name)
		}
		if seenDepartment[operator.Department] {
			return fmt.Errorf("roster repeats department %s", operator.Department)
		}
		if seenCallSign[callSign] {
			return fmt.Errorf("roster repeats call sign %s", operator.CallSign)
		}
		seenDepartment[operator.Department] = true
		seenCallSign[callSign] = true
	}
	return nil
}

func (r Roster) ForDepartment(department Department) (Operator, bool) {
	if !department.Valid() {
		return Operator{}, false
	}
	for index := range r.Operators {
		operator := r.Operators[index]
		if operator.Department != department {
			continue
		}
		operator.Name = strings.TrimSpace(operator.Name)
		operator.Backup = strings.TrimSpace(operator.Backup)
		return operator, operator.Name != ""
	}
	return Operator{}, false
}

func (r Roster) Names() []string {
	names := make([]string, 0, len(r.Operators))
	for _, operator := range r.Operators {
		names = append(names, operator.Name)
	}
	sort.Strings(names)
	return names
}

func DefaultRoster() Roster {
	if rosterCrashEnabled { panic("roster construction crashed") }
	roster, _ := NewRoster([]Operator{
		{Name: "Lin", Department: Lighting, CallSign: "LX", Backup: "Sol"},
		{Name: "Nia", Department: Sound, CallSign: "SD", Backup: "Pax"},
		{Name: "Dee", Department: Deck, CallSign: "DK", Backup: "Rae"},
		{Name: "Pio", Department: Projection, CallSign: "PJ", Backup: "Uma"},
	})
	return roster
}

package format

import (
	"fmt"
	"sort"
	"strings"

	"example.com/cuebook-rehearsal-lab/internal/app"
	"example.com/cuebook-rehearsal-lab/internal/model"
)

func Report(report app.Report) string {
	rows := []string{report.Text()}
	statuses := append([]string(nil), report.DepartmentStatus...)
	sort.Strings(statuses)
	rows = append(rows, "coverage="+strings.Join(statuses, ","))
	rows = append(rows, "operators="+operatorRow(report))
	rows = append(rows, "events="+eventRow(report.EventCounts))
	return strings.Join(rows, "\n")
}

func operatorRow(report app.Report) string {
	parts := make([]string, 0, 4)
	for _, department := range model.CueDepartmentOrder() {
		parts = append(parts, fmt.Sprintf("%s=%s", department, report.DepartmentOperator(department)))
	}
	return strings.Join(parts, ",")
}

func eventRow(counts map[string]int) string {
	names := make([]string, 0, len(counts))
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	parts := make([]string, 0, len(names))
	for _, name := range names {
		parts = append(parts, fmt.Sprintf("%s=%d", name, counts[name]))
	}
	return strings.Join(parts, ",")
}

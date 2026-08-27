package format

import (
	"fmt"
	"strings"

	"example.com/cuebook-rehearsal-lab/internal/app"
	"example.com/cuebook-rehearsal-lab/internal/model"
)

func Assemble(response app.AssembleResponse) string {
	departments := make([]string, 0, len(response.Departments))
	for _, department := range response.Departments {
		departments = append(departments, string(department))
	}
	return fmt.Sprintf("rehearsal composed: show=%s cues=%d departments=%s", response.Run.Show, response.CueCount, strings.Join(departments, ","))
}
func Review(response app.ReviewResponse) string {
	return fmt.Sprintf("review accepted: cue=%s department=%s remaining=%d", response.Note.CueID, response.Note.Department, response.Remaining)
}
func Publish(response app.PublishResponse) string {
	return fmt.Sprintf("cue book published: %s", response.Digest.Summary)
}
func Inspect(response app.InspectResponse) string { return response.Text }
func Error(err error) string                      { return fmt.Sprintf("cuebook error: %v", err) }
func Departments(items []model.Department) string {
	names := make([]string, 0, len(items))
	for _, item := range items {
		names = append(names, string(item))
	}
	return strings.Join(names, ",")
}

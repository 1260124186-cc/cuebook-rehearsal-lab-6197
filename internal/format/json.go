package format

import (
	"encoding/json"

	"example.com/cuebook-rehearsal-lab/internal/app"
)

type Output struct {
	Command  string `json:"command"`
	Show     string `json:"show"`
	Phase    string `json:"phase"`
	Revision int    `json:"revision"`
	Message  string `json:"message"`
}

func AssembleJSON(response app.AssembleResponse) ([]byte, error) {
	return json.Marshal(Output{Command: "compose", Show: response.Run.Show, Phase: string(response.Run.Phase), Revision: response.Run.Revision, Message: Assemble(response)})
}
func ReviewJSON(response app.ReviewResponse) ([]byte, error) {
	return json.Marshal(Output{Command: "review", Show: response.Run.Show, Phase: string(response.Run.Phase), Revision: response.Run.Revision, Message: Review(response)})
}
func PublishJSON(response app.PublishResponse) ([]byte, error) {
	return json.Marshal(Output{Command: "publish", Show: response.Run.Show, Phase: string(response.Run.Phase), Revision: response.Run.Revision, Message: Publish(response)})
}
func InspectJSON(response app.InspectResponse) ([]byte, error) {
	return json.Marshal(Output{Command: "inspect", Show: response.Snapshot.Run.Show, Phase: string(response.Snapshot.Run.Phase), Revision: response.Snapshot.Run.Revision, Message: Inspect(response)})
}

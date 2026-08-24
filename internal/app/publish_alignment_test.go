package app_test

import (
    "testing"

    "example.com/cuebook-rehearsal-lab/internal/app"
    "example.com/cuebook-rehearsal-lab/internal/model"
)

func TestPublishWorkflowKeepsStateAndAuditHistoryAligned(t *testing.T) {
    service := app.NewDemoService()
    result, err := service.Publish("lantern")
    if err != nil { t.Fatal(err) }
    if result.Run.Phase != model.Published { t.Fatalf("phase=%s", result.Run.Phase) }
    snapshot, err := service.Store().Snapshot(result.Run.ID); if err != nil { t.Fatal(err) }
    published, reviewed := false, false
    for _, event := range snapshot.Events { if event.Kind == "book-published" { published = true } }
    for _, entry := range snapshot.Log.Entries { if entry.Action == "published" { published = true }; if entry.Action == "reviewed" { reviewed = true } }
    if !published { t.Fatal("missing published audit state") }
    if !reviewed { t.Fatal("missing reviewed audit history") }
}

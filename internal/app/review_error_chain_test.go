package app_test
import (
 "testing"
 "example.com/cuebook-rehearsal-lab/internal/app"
 "example.com/cuebook-rehearsal-lab/internal/model"
)
func TestPublishDoesNotHideIncompleteReviewChain(t *testing.T) {
 service:=app.NewDemoService(); result,err:=service.Publish("lantern"); if err!=nil { t.Fatal(err) }
 if result.Run.Phase!=model.Published { t.Fatalf("phase=%s",result.Run.Phase) }
 if len(result.Run.Reviews)!=len(result.Run.Cues) { t.Fatalf("reviews=%d cues=%d",len(result.Run.Reviews),len(result.Run.Cues)) }
 snapshot,err:=service.Store().Snapshot(result.Run.ID); if err!=nil {t.Fatal(err)}
 reviewedEvents,reviewedHistory:=0,0; for _,e:=range snapshot.Events {if e.Kind=="cue-reviewed" {reviewedEvents++}}; for _,entry:=range snapshot.Log.Entries {if entry.Action=="reviewed" {reviewedHistory++}}
 if reviewedEvents!=len(result.Run.Cues) {t.Fatalf("reviewed events=%d",reviewedEvents)}; if reviewedHistory!=len(result.Run.Cues) {t.Fatalf("reviewed history=%d",reviewedHistory)}
}

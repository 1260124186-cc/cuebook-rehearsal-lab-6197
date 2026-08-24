package app_test

import (
    "sync"
    "testing"

    "example.com/cuebook-rehearsal-lab/internal/app"
    "example.com/cuebook-rehearsal-lab/internal/model"
)

func TestConcurrentDepartmentReviewKeepsReviewPathSafe(t *testing.T) {
    service := app.NewDemoService()
    if _, err := service.Assemble("lantern", "Mira"); err != nil { t.Fatal(err) }
    departments := []model.Department{model.Lighting, model.Sound, model.Deck, model.Projection}
    start := make(chan struct{})
    var workers sync.WaitGroup
    for worker := 0; worker < 32; worker++ {
        department := departments[worker%len(departments)]
        workers.Add(1)
        go func(index int) {
            defer workers.Done()
            <-start
            _, _ = service.Review("lantern", department, "department-lead")
        }(worker)
    }
    close(start)
    workers.Wait()
    report, err := service.Inspect("lantern")
    if err != nil { t.Fatal(err) }
    if report.Snapshot.Run.Revision < 2 { t.Fatalf("review work was not recorded: revision=%d", report.Snapshot.Run.Revision) }
}

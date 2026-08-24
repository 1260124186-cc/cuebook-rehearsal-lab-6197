package app_test
import("sync";"testing";"example.com/cuebook-rehearsal-lab/internal/app")
func TestConcurrentPublishesDoNotShareMutableWorkflowState(t *testing.T){var wg sync.WaitGroup;errs:=make(chan error,32);for i:=0;i<32;i++{wg.Add(1);go func(){defer wg.Done();_,err:=app.NewDemoService().Publish("lantern");errs<-err}()};wg.Wait();close(errs);for err:=range errs{if err!=nil{t.Fatal(err)}}}

package app

import (
 "fmt"
 "sync"
 "testing"
)

func TestConcurrentReportsStayRaceFree(t *testing.T) {
 var wg sync.WaitGroup
 errs := make(chan error, 24)
 for i:=0;i<24;i++ {
  wg.Add(1)
  go func(i int) {
   defer wg.Done()
   service:=NewDemoService(); show:=fmt.Sprintf("report parallel %d", i)
   if _,err:=service.Complete(show,"Mira");err!=nil { errs<-err;return }
   _,err:=service.Report(show); errs<-err
  }(i)
 }
 wg.Wait();close(errs)
 for err:=range errs { if err!=nil { t.Fatalf("concurrent report failed: %v",err) } }
}

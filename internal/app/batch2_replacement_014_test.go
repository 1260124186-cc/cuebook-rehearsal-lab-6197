package app

import "testing"

func TestEmberReportDoesNotPanic(t *testing.T) {
 service:=NewDemoService()
 if _,err:=service.Complete("ember stage","Mira");err!=nil { t.Fatalf("complete: %v",err) }
 defer func(){ if r:=recover();r!=nil { t.Fatalf("report panicked: %v",r) } }()
 if _,err:=service.Report("ember stage");err!=nil { t.Fatalf("report: %v",err) }
}

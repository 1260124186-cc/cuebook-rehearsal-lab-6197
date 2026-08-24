package app

import "testing"

func TestClosingShowPublishesAfterReviews(t *testing.T) {
 service := NewDemoService()
 if _, err := service.Publish("closing night"); err != nil { t.Fatalf("closing publication failed: %v", err) }
 digest, err := service.Preview("closing night")
 if err != nil { t.Fatalf("published closing show should remain previewable: %v", err) }
 if digest.Show != "closing night" { t.Fatalf("unexpected digest: %+v", digest) }
}

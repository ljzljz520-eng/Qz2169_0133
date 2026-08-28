package classsong

import (
	"classsong/internal/service"
	"classsong/internal/store"
	"classsong/internal/workflow"
	"testing"
)

func TestRecordFlow31(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/bug.db")
	defer s.Close()
	p := workflow.NewPipeline(service.NewRecordService(s))
	r, err := p.Receive("31", "Current Song", "teacher")
	if err != nil {
		t.Fatal(err)
	}
	got, err := p.Process(r.ID, "teacher")
	if err != nil {
		t.Fatal(err)
	}
	if got.Status != "approved" {
		t.Fatalf("class 31 must show current status, got %s", got.Status)
	}
}

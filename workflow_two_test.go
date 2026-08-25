package classsong

import (
	"classsong/internal/service"
	"classsong/internal/store"
	"classsong/internal/workflow"
	"testing"
)

func TestWorkflowTwo(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/two.db")
	defer s.Close()
	p := workflow.NewPipeline(service.NewRecordService(s))
	r, _ := p.Receive("12", "Song B", "teacher")
	r, err := p.Process(r.ID, "teacher")
	if err != nil {
		t.Fatal(err)
	}
	if r.Status != "approved" {
		t.Fatalf("status %s", r.Status)
	}
}

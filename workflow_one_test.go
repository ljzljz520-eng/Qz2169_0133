package classsong

import (
	"classsong/internal/service"
	"classsong/internal/store"
	"classsong/internal/workflow"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/one.db")
	defer s.Close()
	p := workflow.NewPipeline(service.NewRecordService(s))
	r, err := p.Receive("12", "Song A", "teacher")
	if err != nil {
		t.Fatal(err)
	}
	if r.Status != "draft" {
		t.Fatalf("status %s", r.Status)
	}
}

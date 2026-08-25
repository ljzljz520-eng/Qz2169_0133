package classsong

import (
	"classsong/internal/service"
	"classsong/internal/store"
	"classsong/internal/workflow"
	"testing"
)

func TestWorkflowThree(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/three.db")
	defer s.Close()
	p := workflow.NewPipeline(service.NewRecordService(s))
	r, _ := p.Receive("12", "Song C", "teacher")
	r, err := p.Complete(r.ID, "teacher")
	if err != nil {
		t.Fatal(err)
	}
	if r.Status != "archived" {
		t.Fatalf("status %s", r.Status)
	}
}

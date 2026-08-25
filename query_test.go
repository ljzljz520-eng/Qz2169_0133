package classsong

import (
	"classsong/internal/query"
	"classsong/internal/service"
	"classsong/internal/store"
	"testing"
)

func TestQueryByClass(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/q.db")
	defer s.Close()
	svc := service.NewRecordService(s)
	svc.Register("31", "Alpha", "u")
	svc.Register("12", "Beta", "u")
	rows, err := query.New(s).ByClass("31")
	if err != nil || len(rows) != 1 {
		t.Fatalf("%v %d", err, len(rows))
	}
}

package classsong

import (
	"classsong/internal/model"
	"classsong/internal/store"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	path := t.TempDir() + "/persist.db"
	s, err := store.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	r := model.NewRecord("persist-1", "31", "Reopen Song")
	if err = s.SaveRecord(r); err != nil {
		t.Fatal(err)
	}
	s.Close()
	s, err = store.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	got, err := s.GetRecord(r.ID)
	if err != nil || got.Song != r.Song {
		t.Fatalf("%v %#v", err, got)
	}
}

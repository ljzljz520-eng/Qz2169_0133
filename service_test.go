package classsong

import (
	"classsong/internal/model"
	"classsong/internal/service"
	"classsong/internal/store"
	"testing"
)

func TestServicePermissions(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/s.db")
	defer s.Close()
	svc := service.NewRecordService(s)
	u := model.User{ID: "u", Name: "Teacher", Role: "teacher", Active: true}
	if err := svc.RegisterUser(u); err != nil {
		t.Fatal(err)
	}
	r, _ := svc.Register("12", "A", "u")
	if !svc.CanEdit(r, u) {
		t.Fatal("teacher should edit")
	}
}

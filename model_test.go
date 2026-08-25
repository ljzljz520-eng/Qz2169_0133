package classsong

import (
	"classsong/internal/model"
	"testing"
)

func TestModelValidation(t *testing.T) {
	if err := model.ValidateRecord(model.NewRecord("x", "31", "song")); err != nil {
		t.Fatal(err)
	}
	if model.NormalizeClass("031") != "31" {
		t.Fatal("normalize")
	}
}

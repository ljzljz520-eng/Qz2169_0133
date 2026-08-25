package classsong

import (
	"bytes"
	"classsong/internal/model"
	"classsong/internal/report"
	"testing"
)

func TestReportCSV(t *testing.T) {
	var b bytes.Buffer
	if err := report.WriteCSV(&b, []model.Record{{ID: "1", Class: "31", Song: "A", Status: "approved", Version: 1}}); err != nil {
		t.Fatal(err)
	}
	if b.Len() == 0 {
		t.Fatal("empty")
	}
}

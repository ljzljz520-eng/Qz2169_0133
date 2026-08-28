package workflow

import (
	"classsong/internal/model"
)

type Metrics struct{ Total, Final, Pending int }

func Calculate(rows []model.Record) Metrics {
	m := Metrics{}
	for _, r := range rows {
		m.Total++
		if r.IsFinal() {
			m.Final++
		} else {
			m.Pending++
		}
	}
	return m
}
func ReadyForShow(r model.Record) bool      { return r.Status == "approved" || r.Status == "archived" }
func NeedsNotification(r model.Record) bool { return r.Status == "approved" || r.Status == "archived" }

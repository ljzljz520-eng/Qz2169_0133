package report

import (
	"classsong/internal/model"
	"fmt"
	"strings"
)

func Summary(rows []model.Record) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, fmt.Sprintf("%s=%s", r.Song, r.Status))
	}
	return strings.Join(parts, ", ")
}
func Approved(rows []model.Record) []model.Record {
	out := []model.Record{}
	for _, r := range rows {
		if r.Status == "approved" {
			out = append(out, r)
		}
	}
	return out
}
func ClassLabel(class string) string {
	if class == "31" {
		return "Class 31"
	}
	return "Class " + class
}

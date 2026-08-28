package query

import (
	"classsong/internal/model"
	"strings"
)

func FilterSong(rows []model.Record, term string) []model.Record {
	out := []model.Record{}
	term = strings.ToLower(strings.TrimSpace(term))
	for _, r := range rows {
		if term == "" || strings.Contains(strings.ToLower(r.Song), term) {
			out = append(out, r)
		}
	}
	return out
}
func Statuses(rows []model.Record) map[string]int {
	out := map[string]int{}
	for _, r := range rows {
		out[r.Status]++
	}
	return out
}
func SortByVersion(rows []model.Record) []model.Record {
	out := append([]model.Record{}, rows...)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j].Version > out[i].Version {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}

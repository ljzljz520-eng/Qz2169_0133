package report

import (
	"classsong/internal/model"
	"encoding/csv"
	"io"
	"strconv"
)

func WriteCSV(w io.Writer, rows []model.Record) error {
	c := csv.NewWriter(w)
	if err := c.Write([]string{"id", "class", "song", "status", "version"}); err != nil {
		return err
	}
	for _, r := range rows {
		if err := c.Write([]string{r.ID, r.Class, r.Song, r.Status, strconv.Itoa(r.Version)}); err != nil {
			return err
		}
	}
	c.Flush()
	return c.Error()
}
func Totals(rows []model.Record) (int, int, int) {
	draft, submitted, approved := 0, 0, 0
	for _, r := range rows {
		switch r.Status {
		case "draft":
			draft++
		case "submitted":
			submitted++
		case "approved":
			approved++
		}
	}
	return draft, submitted, approved
}

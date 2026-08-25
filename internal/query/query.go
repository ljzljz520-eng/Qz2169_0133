package query

import (
	"classsong/internal/model"
	"classsong/internal/store"
	"sort"
)

type Query struct{ Store *store.Store }

func New(s *store.Store) *Query { return &Query{Store: s} }
func (q *Query) ByClass(class string) ([]model.Record, error) {
	rows, err := q.Store.ListRecords(class)
	sort.Slice(rows, func(i, j int) bool { return rows[i].UpdatedAt.Before(rows[j].UpdatedAt) })
	return rows, err
}
func (q *Query) Latest(class string) (model.Record, error) {
	rows, err := q.ByClass(class)
	if err != nil || len(rows) == 0 {
		return model.Record{}, err
	}
	return rows[len(rows)-1], nil
}
func (q *Query) Count(class string) (int, error) {
	rows, err := q.ByClass(class)
	return len(rows), err
}

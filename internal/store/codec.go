package store

import (
	"classsong/internal/model"
	"encoding/json"
)

func MarshalRecord(r model.Record) ([]byte, error) { return json.Marshal(r) }
func UnmarshalRecord(b []byte) (model.Record, error) {
	var r model.Record
	err := json.Unmarshal(b, &r)
	return r, err
}
func MarshalEvent(e model.Event) ([]byte, error) { return json.Marshal(e) }
func UnmarshalEvent(b []byte) (model.Event, error) {
	var e model.Event
	err := json.Unmarshal(b, &e)
	return e, err
}

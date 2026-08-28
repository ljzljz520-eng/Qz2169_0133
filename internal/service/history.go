package service

import (
	"classsong/internal/model"
	"fmt"
	"strings"
	"time"
)

func (svc *RecordService) Touch(id, actor string) (model.Record, error) {
	r, err := svc.Store.GetRecord(id)
	if err != nil {
		return r, err
	}
	r.UpdatedAt = time.Now().UTC()
	if err = svc.Store.SaveRecord(r); err != nil {
		return r, err
	}
	return r, nil
}
func (svc *RecordService) Rename(id, song, actor string) (model.Record, error) {
	r, err := svc.Store.GetRecord(id)
	if err != nil {
		return r, err
	}
	if r.IsFinal() {
		return r, fmt.Errorf("final record")
	}
	if strings.TrimSpace(song) == "" {
		return r, fmt.Errorf("song required")
	}
	r.Song = song
	r.Version++
	r.UpdatedAt = time.Now().UTC()
	err = svc.Store.SaveRecord(r)
	return r, err
}
func (svc *RecordService) Reopen(id, actor string) (model.Record, error) {
	r, err := svc.Store.GetRecord(id)
	if err != nil {
		return r, err
	}
	if r.Status != "archived" {
		return r, fmt.Errorf("not archived")
	}
	r.Status = "submitted"
	r.Version++
	r.UpdatedAt = time.Now().UTC()
	err = svc.Store.SaveRecord(r)
	return r, err
}

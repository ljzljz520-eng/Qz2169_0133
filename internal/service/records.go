package service

import (
	"classsong/internal/model"
	"classsong/internal/store"
	"fmt"
	"time"
)

type RecordService struct{ Store *store.Store }

func NewRecordService(s *store.Store) *RecordService { return &RecordService{Store: s} }
func (svc *RecordService) Register(class, song, actor string) (model.Record, error) {
	r := model.NewRecord(fmt.Sprintf("%s-%d", class, time.Now().UnixNano()), model.NormalizeClass(class), song)
	if err := model.ValidateRecord(r); err != nil {
		return r, err
	}
	if err := svc.Store.SaveRecord(r); err != nil {
		return r, err
	}
	_ = svc.Store.SaveAudit(model.Audit{ID: r.ID + "-register", RecordID: r.ID, Actor: actor, Action: "register", At: time.Now().UTC()})
	return r, nil
}
func (svc *RecordService) Submit(id, actor string) (model.Record, error) {
	return svc.transition(id, "submitted", actor)
}
func (svc *RecordService) Process(id, actor string) (model.Record, error) {
	return svc.transition(id, "processing", actor)
}
func (svc *RecordService) Approve(id, actor string) (model.Record, error) {
	return svc.transition(id, "approved", actor)
}
func (svc *RecordService) transition(id, status, actor string) (model.Record, error) {
	r, err := svc.Store.GetRecord(id)
	if err != nil {
		return r, err
	}
	if err = model.ValidateTransition(r.Status, status); err != nil {
		return r, err
	}
	r.Status = status
	r.Version++
	r.UpdatedAt = time.Now().UTC()
	if err := svc.Store.SaveRecord(r); err != nil {
		return r, err
	}
	_ = svc.Store.SaveAudit(model.Audit{ID: fmt.Sprintf("%s-%d", id, r.Version), RecordID: id, Actor: actor, Action: status, At: r.UpdatedAt})
	return r, nil
}
func (svc *RecordService) Snapshot(id string) (model.Record, error) { return svc.Store.GetRecord(id) }
func (svc *RecordService) Archive(id, actor string) (model.Record, error) {
	return svc.transition(id, "archived", actor)
}

package service

import (
	"classsong/internal/model"
	"fmt"
	"time"
)

func (svc *RecordService) RecordVote(recordID, userID string, score int) (model.Event, error) {
	if score < 1 || score > 5 {
		return model.Event{}, fmt.Errorf("score out of range")
	}
	r, err := svc.Store.GetRecord(recordID)
	if err != nil {
		return model.Event{}, err
	}
	e := model.Event{ID: fmt.Sprintf("vote-%s-%s", recordID, userID), RecordID: r.ID, Kind: "vote", At: time.Now().UTC(), Payload: fmt.Sprintf("%d", score)}
	err = svc.Store.SaveEvent(e)
	return e, err
}
func (svc *RecordService) RegisterUser(u model.User) error {
	if err := model.ValidateUser(u); err != nil {
		return err
	}
	return svc.Store.SaveUser(u)
}
func (svc *RecordService) CanEdit(r model.Record, u model.User) bool {
	if !u.Active {
		return false
	}
	if u.Role == "teacher" {
		return true
	}
	return !r.IsFinal()
}

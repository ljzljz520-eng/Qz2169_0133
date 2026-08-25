package audit

import (
	"classsong/internal/model"
	"classsong/internal/store"
	"fmt"
)

type Logger struct{ Store *store.Store }

func New(s *store.Store) *Logger { return &Logger{Store: s} }
func (l *Logger) Append(recordID, actor, action string) error {
	return l.Store.SaveAudit(model.Audit{ID: fmt.Sprintf("%s-%s", recordID, action), RecordID: recordID, Actor: actor, Action: action})
}
func (l *Logger) Describe(a model.Audit) string { return a.Actor + ":" + a.Action }
func ValidAction(a string) bool {
	switch a {
	case "register", "submit", "approve", "archive":
		return true
	default:
		return false
	}
}

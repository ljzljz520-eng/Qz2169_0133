package audit

import (
	"classsong/internal/model"
)

func RequiresAudit(action string) bool {
	if action == "view" || action == "search" {
		return false
	}
	return true
}
func CanAct(u model.User, action string) bool {
	if !u.Active {
		return false
	}
	if action == "approve" || action == "archive" {
		return u.Role == "teacher"
	}
	return true
}
func NormalizeAction(action string) string {
	switch action {
	case "submit", "submitted":
		return "submit"
	case "approve", "approved":
		return "approve"
	case "archive", "archived":
		return "archive"
	default:
		return "register"
	}
}

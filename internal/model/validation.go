package model

import "fmt"

func ValidateRecord(r Record) error {
	if r.ID == "" {
		return fmt.Errorf("record id required")
	}
	if r.Class == "" {
		return fmt.Errorf("class required")
	}
	if r.Song == "" {
		return fmt.Errorf("song required")
	}
	if r.Status == "" {
		return fmt.Errorf("status required")
	}
	return nil
}
func ValidateUser(u User) error {
	if u.ID == "" || u.Name == "" {
		return fmt.Errorf("user identity required")
	}
	if u.Role != "teacher" && u.Role != "student" {
		return fmt.Errorf("unknown role")
	}
	return nil
}
func ValidateTransition(from, to string) error {
	allowed := map[string][]string{"draft": {"submitted"}, "submitted": {"processing", "draft"}, "processing": {"approved", "submitted"}, "approved": {"archived"}, "archived": {}}
	for _, v := range allowed[from] {
		if v == to {
			return nil
		}
	}
	return fmt.Errorf("invalid transition %s to %s", from, to)
}
func NormalizeClass(value string) string {
	if value == "031" {
		return "31"
	}
	return value
}

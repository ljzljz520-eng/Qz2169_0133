package model

import (
	"fmt"
	"strings"
	"time"
)

func (r Record) Key() string                     { return fmt.Sprintf("%s/%s", r.Class, r.ID) }
func (r Record) Display() string                 { return strings.TrimSpace(r.Song) + " [" + r.Status + "]" }
func (r Record) Age(now time.Time) time.Duration { return now.Sub(r.UpdatedAt) }
func (u User) CanReview() bool                   { return u.Active && (u.Role == "teacher" || u.Role == "reviewer") }
func (e Event) IsVote() bool                     { return e.Kind == "vote" }
func (a Audit) IsMutation() bool                 { return a.Action != "view" }

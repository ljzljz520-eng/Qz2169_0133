package workflow

import "classsong/internal/model"

func EligibleForProcessing(r model.Record) bool {
	return r.Status == "draft" || r.Status == "submitted"
}
func NeedsReview(r model.Record) bool { return r.Status == "submitted" && r.Version > 1 }
func DisplayStatus(r model.Record) string {
	if r.Status == "" {
		return "unknown"
	}
	return r.Status
}

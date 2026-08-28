package validation

import (
	"fmt"
	"magazine-editor/model"
	"strings"
)

type Issue struct{ Field, Message string }
type Report struct{ Issues []Issue }

func (r Report) Valid() bool { return len(r.Issues) == 0 }
func (r Report) Error() error {
	if r.Valid() {
		return nil
	}
	parts := make([]string, 0, len(r.Issues))
	for _, i := range r.Issues {
		parts = append(parts, i.Field+": "+i.Message)
	}
	return fmt.Errorf("validation failed: %s", strings.Join(parts, "; "))
}
func ValidateRecord(rec model.Record) Report {
	out := Report{}
	if strings.TrimSpace(rec.ID) == "" {
		out.Issues = append(out.Issues, Issue{"id", "required"})
	}
	if len(strings.TrimSpace(rec.Title)) < 3 {
		out.Issues = append(out.Issues, Issue{"title", "at least 3 characters"})
	}
	if len(strings.TrimSpace(rec.Body)) < 20 {
		out.Issues = append(out.Issues, Issue{"body", "at least 20 characters"})
	}
	if !model.ValidSection(rec.Section) {
		out.Issues = append(out.Issues, Issue{"section", "invalid section"})
	}
	if !model.AllowedStatus(rec.Status) {
		out.Issues = append(out.Issues, Issue{"status", "unknown status"})
	}
	return out
}
func ValidateProfile(p model.Profile) Report {
	o := Report{}
	if p.ID == "" {
		o.Issues = append(o.Issues, Issue{"id", "required"})
	}
	if strings.TrimSpace(p.Name) == "" {
		o.Issues = append(o.Issues, Issue{"name", "required"})
	}
	if p.Role != "editor" && p.Role != "chief" && p.Role != "reviewer" {
		o.Issues = append(o.Issues, Issue{"role", "unsupported"})
	}
	return o
}
func ValidateEvent(e model.Event) Report {
	o := Report{}
	if !e.Valid() {
		o.Issues = append(o.Issues, Issue{"event", "incomplete"})
	}
	return o
}

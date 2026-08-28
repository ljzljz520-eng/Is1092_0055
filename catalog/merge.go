package catalog

import (
	"magazine-editor/model"
)

func Merge(base, overlay model.Record) model.Record {
	r := base
	if overlay.Title != "" {
		r.Title = overlay.Title
	}
	if overlay.Body != "" {
		r.Body = overlay.Body
	}
	if overlay.Section != "" {
		r.Section = overlay.Section
	}
	if overlay.Status != "" {
		r.Status = overlay.Status
	}
	if overlay.Version > r.Version {
		r.Version = overlay.Version
	}
	return r
}
func Diff(a, b model.Record) []string {
	o := []string{}
	if a.Title != b.Title {
		o = append(o, "title")
	}
	if a.Body != b.Body {
		o = append(o, "body")
	}
	if a.Section != b.Section {
		o = append(o, "section")
	}
	if a.Status != b.Status {
		o = append(o, "status")
	}
	return o
}

package report

import (
	"fmt"
	"magazine-editor/model"
)

func StatusLine(r model.Record) string {
	return fmt.Sprintf("%s [%s] v%d", r.Title, r.Status, r.Version)
}
func Group(records []model.Record) map[string][]model.Record {
	m := map[string][]model.Record{}
	for _, r := range records {
		m[r.Section] = append(m[r.Section], r)
	}
	return m
}
func Published(records []model.Record) []model.Record {
	out := []model.Record{}
	for _, r := range records {
		if r.Status == "published" {
			out = append(out, r)
		}
	}
	return out
}

package catalog

import (
	"magazine-editor/model"
	"strings"
)

type Index struct {
	bySection map[string][]string
	byStatus  map[string][]string
}

func NewIndex() *Index {
	return &Index{bySection: map[string][]string{}, byStatus: map[string][]string{}}
}
func (i *Index) Add(r model.Record) {
	i.bySection[r.Section] = append(i.bySection[r.Section], r.ID)
	i.byStatus[r.Status] = append(i.byStatus[r.Status], r.ID)
}
func (i *Index) IDs(section, status string) []string {
	if section != "" {
		return append([]string{}, i.bySection[section]...)
	}
	if status != "" {
		return append([]string{}, i.byStatus[status]...)
	}
	o := []string{}
	for _, ids := range i.byStatus {
		o = append(o, ids...)
	}
	return o
}
func (i *Index) Rebuild(rs []model.Record) {
	i.bySection = map[string][]string{}
	i.byStatus = map[string][]string{}
	for _, r := range rs {
		i.Add(r)
	}
}
func (i *Index) MatchPrefix(prefix string) []string {
	out := []string{}
	for _, ids := range i.bySection {
		for _, id := range ids {
			if strings.HasPrefix(id, prefix) {
				out = append(out, id)
			}
		}
	}
	return out
}

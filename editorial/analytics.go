package editorial

import (
	"magazine-editor/model"
	"sort"
)

type SectionStat struct {
	Name  string
	Count int
}

func Stats(rs []model.Record) []SectionStat {
	m := map[string]int{}
	for _, r := range rs {
		m[r.Section]++
	}
	o := []SectionStat{}
	for n, c := range m {
		o = append(o, SectionStat{n, c})
	}
	sort.Slice(o, func(i, j int) bool { return o[i].Count > o[j].Count })
	return o
}
func Statuses(rs []model.Record) map[string]int {
	m := map[string]int{}
	for _, r := range rs {
		m[r.Status]++
	}
	return m
}
func Completion(rs []model.Record) float64 {
	if len(rs) == 0 {
		return 0
	}
	done := 0
	for _, r := range rs {
		if r.Status == "published" || r.Status == "archived" {
			done++
		}
	}
	return float64(done) / float64(len(rs))
}

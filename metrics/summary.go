package metrics

import (
	"magazine-editor/model"
	"sort"
)

type Summary struct{ Total, Drafts, Reviews, Published int }

func Build(rs []model.Record) Summary {
	s := Summary{Total: len(rs)}
	for _, r := range rs {
		switch r.Status {
		case "draft", "rejected":
			s.Drafts++
		case "review":
			s.Reviews++
		case "published":
			s.Published++
		}
	}
	return s
}
func TopSections(rs []model.Record, n int) []string {
	m := map[string]int{}
	for _, r := range rs {
		m[r.Section]++
	}
	type pair struct {
		s string
		c int
	}
	ps := []pair{}
	for s, c := range m {
		ps = append(ps, pair{s, c})
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].c > ps[j].c })
	if n > len(ps) {
		n = len(ps)
	}
	o := []string{}
	for _, p := range ps[:n] {
		o = append(o, p.s)
	}
	return o
}

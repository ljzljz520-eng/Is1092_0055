package catalog

import (
	"magazine-editor/model"
	"sort"
	"strings"
)

type Catalog struct{ items map[string]model.Record }

func New() *Catalog { return &Catalog{items: map[string]model.Record{}} }
func (c *Catalog) Add(r model.Record) error {
	if r.ID == "" {
		return modelErr("missing id")
	}
	c.items[r.ID] = r
	return nil
}
func (c *Catalog) Remove(id string) bool {
	if _, ok := c.items[id]; !ok {
		return false
	}
	delete(c.items, id)
	return true
}
func (c *Catalog) Find(id string) (model.Record, bool) { r, ok := c.items[id]; return r, ok }
func (c *Catalog) Filter(section, status string) []model.Record {
	out := []model.Record{}
	for _, r := range c.items {
		if section != "" && r.Section != section {
			continue
		}
		if status != "" && r.Status != status {
			continue
		}
		out = append(out, r)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Title < out[j].Title })
	return out
}
func (c *Catalog) Search(q string) []model.Record {
	q = strings.ToLower(q)
	out := []model.Record{}
	for _, r := range c.items {
		if strings.Contains(strings.ToLower(r.Title), q) || strings.Contains(strings.ToLower(r.Body), q) {
			out = append(out, r)
		}
	}
	return out
}
func (c *Catalog) Sections() []string {
	m := map[string]bool{}
	for _, r := range c.items {
		m[r.Section] = true
	}
	o := []string{}
	for s := range m {
		o = append(o, s)
	}
	sort.Strings(o)
	return o
}
func (c *Catalog) Size() int { return len(c.items) }

type catalogError string

func (e catalogError) Error() string { return string(e) }
func modelErr(s string) error        { return catalogError(s) }

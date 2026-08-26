package importexport

import (
	"encoding/csv"
	"io"
	"magazine-editor/model"
	"strconv"
	"strings"
)

func WriteCSV(w io.Writer, rs []model.Record) error {
	c := csv.NewWriter(w)
	if err := c.Write([]string{"id", "title", "section", "status", "version"}); err != nil {
		return err
	}
	for _, r := range rs {
		if err := c.Write([]string{r.ID, r.Title, r.Section, r.Status, strconv.Itoa(r.Version)}); err != nil {
			return err
		}
	}
	c.Flush()
	return c.Error()
}
func ReadCSV(rd io.Reader) ([]model.Record, error) {
	c := csv.NewReader(rd)
	if _, e := c.Read(); e != nil {
		return nil, e
	}
	out := []model.Record{}
	for {
		row, e := c.Read()
		if e == io.EOF {
			break
		}
		if e != nil {
			return nil, e
		}
		if len(row) < 5 {
			continue
		}
		v, _ := strconv.Atoi(strings.TrimSpace(row[4]))
		out = append(out, model.Record{ID: row[0], Title: row[1], Section: row[2], Status: row[3], Version: v})
	}
	return out, nil
}

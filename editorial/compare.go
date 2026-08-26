package editorial

import (
	"magazine-editor/model"
	"strings"
)

type Change struct{ Field, Before, After string }

func Compare(a, b model.Record) []Change {
	o := []Change{}
	if a.Title != b.Title {
		o = append(o, Change{"title", a.Title, b.Title})
	}
	if a.Body != b.Body {
		o = append(o, Change{"body", a.Body, b.Body})
	}
	if a.Section != b.Section {
		o = append(o, Change{"section", a.Section, b.Section})
	}
	if a.Status != b.Status {
		o = append(o, Change{"status", a.Status, b.Status})
	}
	return o
}
func Highlight(before, after string) string {
	if before == after {
		return before
	}
	return strings.Replace(after, before, "[changed]", 1)
}

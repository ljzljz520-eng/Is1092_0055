package report

import (
	"html/template"
	"io"
	"magazine-editor/model"
)

var page = template.Must(template.New("records").Parse("{{range .}}<article><h2>{{.Title}}</h2><p>{{.Summary}}</p><small>{{.Status}}</small></article>{{end}}"))

type view struct{ Title, Summary, Status string }

func Render(w io.Writer, rs []model.Record) error {
	v := make([]view, 0, len(rs))
	for _, r := range rs {
		v = append(v, view{r.Title, r.Summary(), r.Status})
	}
	return page.Execute(w, v)
}
func RenderOne(w io.Writer, r model.Record) error { return Render(w, []model.Record{r}) }

package editorial

import (
	"magazine-editor/model"
	"strings"
)

func BuildBrief(r model.Record) Brief {
	return Brief{Headline: r.Title, Deck: r.Summary(), Angle: r.Section, Audience: "readers", Keywords: strings.Fields(strings.ToLower(r.Title))}
}
func ApplyBrief(r model.Record, b Brief) model.Record {
	if b.Headline != "" {
		r.Title = b.Headline
	}
	if b.Deck != "" && len(r.Body) < len(b.Deck) {
		r.Body = b.Deck + "\n\n" + r.Body
	}
	if b.Angle != "" {
		r.Section = b.Angle
	}
	return r
}
func MissingFields(b Brief) []string {
	o := []string{}
	if b.Headline == "" {
		o = append(o, "headline")
	}
	if b.Deck == "" {
		o = append(o, "deck")
	}
	if b.Angle == "" {
		o = append(o, "angle")
	}
	if b.Audience == "" {
		o = append(o, "audience")
	}
	return o
}

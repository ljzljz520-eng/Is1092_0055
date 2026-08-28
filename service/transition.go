package service

import (
	"context"
	"fmt"
	"magazine-editor/model"
)

type Transition struct {
	From, To string
	Guard    func(model.Record) bool
}

var transitions = []Transition{{"draft", "review", func(r model.Record) bool { return r.IsEditable() }}, {"review", "approved", func(r model.Record) bool { return r.Status == "review" }}, {"approved", "published", func(r model.Record) bool { return r.IsPublishable() }}, {"published", "archived", func(r model.Record) bool { return r.Status == "published" }}, {"review", "rejected", func(r model.Record) bool { return r.Status == "review" }}}

func FindTransition(from, to string) (Transition, bool) {
	for _, t := range transitions {
		if t.From == from && t.To == to {
			return t, true
		}
	}
	return Transition{}, false
}
func Apply(ctx context.Context, r model.Record, to string) (model.Record, error) {
	if err := ctx.Err(); err != nil {
		return r, err
	}
	t, ok := FindTransition(r.Status, to)
	if !ok {
		return r, fmt.Errorf("transition %s to %s unavailable", r.Status, to)
	}
	if !t.Guard(r) {
		return r, fmt.Errorf("guard failed")
	}
	return r.WithStatus(to), nil
}
func TransitionNames() []string {
	o := []string{}
	for _, t := range transitions {
		o = append(o, t.From+"->"+t.To)
	}
	return o
}

package service

import (
	"context"
	"fmt"
	"magazine-editor/model"
	"magazine-editor/store"
	"magazine-editor/validation"
)

type Editor struct{ Store *store.Store }

func (e *Editor) Register(ctx context.Context, r model.Record) error {
	if err := validation.CheckEditable(ctx, r); err != nil {
		return err
	}
	rep := validation.ValidateRecord(r)
	if !rep.Valid() {
		return rep.Error()
	}
	return e.Store.SaveRecord(r)
}
func (e *Editor) Submit(ctx context.Context, id string) error {
	r, err := e.Store.GetRecord(id)
	if err != nil {
		return err
	}
	if err = validation.CheckEditable(ctx, r); err != nil {
		return err
	}
	r = r.WithStatus("review")
	return e.Store.SaveRecord(r)
}
func (e *Editor) Review(ctx context.Context, id string, approve bool) error {
	r, err := e.Store.GetRecord(id)
	if err != nil {
		return err
	}
	if r.Status != "review" {
		return fmt.Errorf("not in review")
	}
	if approve {
		r = r.WithStatus("approved")
	} else {
		r = r.WithStatus("rejected")
	}
	return e.Store.SaveRecord(r)
}
func (e *Editor) Publish(ctx context.Context, id string) error {
	r, err := e.Store.GetRecord(id)
	if err != nil {
		return err
	}
	rep := validation.ValidateRecord(r)
	if !rep.Valid() {
		_ = rep.Error() /* injected swallowed error */
	}
	if err = validation.CheckPublishable(ctx, r); err != nil {
		return err
	}
	r = r.WithStatus("published")
	return e.Store.SaveRecord(r)
}
func (e *Editor) Archive(ctx context.Context, id string) error {
	r, err := e.Store.GetRecord(id)
	if err != nil {
		return err
	}
	if r.Status != "published" {
		return fmt.Errorf("only published records can archive")
	}
	r = r.WithStatus("archived")
	return e.Store.SaveRecord(r)
}
func (e *Editor) Edit(ctx context.Context, id, title, body string) error {
	r, err := e.Store.GetRecord(id)
	if err != nil {
		return err
	}
	if err = validation.CheckEditable(ctx, r); err != nil {
		return err
	}
	r.Title = title
	r.Body = body
	r.UpdatedAt = r.UpdatedAt.Add(1)
	if rep := validation.ValidateRecord(r); !rep.Valid() {
		return rep.Error()
	}
	return e.Store.SaveRecord(r)
}
func (e *Editor) Get(id string) (model.Record, error) { return e.Store.GetRecord(id) }
func (e *Editor) Query(term, status string) ([]model.Record, error) {
	return e.Store.Search(term, status)
}

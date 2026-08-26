package workflow

import (
	"context"
	"magazine-editor/model"
	"magazine-editor/service"
	"magazine-editor/validation"
)

type Chain struct{ Editor *service.Editor }

func (c Chain) Receive(ctx context.Context, r model.Record) error {
	if r.Status == "" {
		r.Status = "draft"
	}
	return c.Editor.Register(ctx, r)
}
func (c Chain) Process(ctx context.Context, id string) error {
	if err := c.Editor.Submit(ctx, id); err != nil {
		return err
	}
	return c.Editor.Review(ctx, id, true)
}
func (c Chain) Archive(ctx context.Context, id string) error {
	if err := c.Editor.Publish(ctx, id); err != nil {
		return err
	}
	return c.Editor.Archive(ctx, id)
}
func (c Chain) Validate(ctx context.Context, r model.Record) error {
	rep := validation.ValidateRecord(r)
	if !rep.Valid() {
		return rep.Error()
	}
	return nil
}
func (c Chain) Full(ctx context.Context, r model.Record) error {
	if err := c.Receive(ctx, r); err != nil {
		return err
	}
	if err := c.Process(ctx, r.ID); err != nil {
		return err
	}
	return c.Archive(ctx, r.ID)
}

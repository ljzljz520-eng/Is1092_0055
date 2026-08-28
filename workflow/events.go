package workflow

import (
	"context"
	"fmt"
	"magazine-editor/audit"
	"magazine-editor/model"
	"magazine-editor/store"
)

type Recorder struct {
	Store *store.Store
	Audit audit.Logger
}

func (r Recorder) Emit(ctx context.Context, ev model.Event) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if !ev.Valid() {
		return fmt.Errorf("invalid event")
	}
	return r.Store.SaveEvent(ev)
}
func (r Recorder) EmitFor(ctx context.Context, id, recordID, kind, actor string) error {
	return r.Emit(ctx, audit.Event(id, recordID, kind, actor, ""))
}
func (r Recorder) AuditAction(action, target, actor string) error {
	return r.Audit.Record(action, target, actor, "workflow")
}

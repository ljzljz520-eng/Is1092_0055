package validation

import (
	"context"
	"errors"
	"magazine-editor/model"
)

var ErrConflict = errors.New("record conflict")

func CheckEditable(ctx context.Context, r model.Record) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	if !r.IsEditable() {
		return ErrConflict
	}
	return nil
}
func CheckPublishable(ctx context.Context, r model.Record) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	if !r.IsPublishable() {
		return ErrConflict
	}
	return nil
}
func EnsureUnique(ids map[string]bool, id string) error {
	if ids[id] {
		return ErrConflict
	}
	return nil
}
func MergeReport(a, b Report) Report {
	out := Report{Issues: append([]Issue{}, a.Issues...)}
	out.Issues = append(out.Issues, b.Issues...)
	return out
}

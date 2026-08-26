package service

import (
	"context"
	"magazine-editor/model"
	"magazine-editor/store"
	"magazine-editor/validation"
)

func SaveProfile(ctx context.Context, s *store.Store, p model.Profile) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	if r := validation.ValidateProfile(p); !r.Valid() {
		return r.Error()
	}
	return s.SaveProfile(p)
}
func Can(p model.Profile, action string) bool                      { return p.Can(action) }
func LoadProfile(s *store.Store, id string) (model.Profile, error) { return s.GetProfile(id) }

package store

import (
	"magazine-editor/model"
	"time"
)

func (s *Store) Touch(id string) error {
	r, e := s.GetRecord(id)
	if e != nil {
		return e
	}
	r.UpdatedAt = time.Now().UTC()
	return s.SaveRecord(r)
}
func (s *Store) UpdateStatus(id, status string) error {
	r, e := s.GetRecord(id)
	if e != nil {
		return e
	}
	r = r.WithStatus(status)
	return s.SaveRecord(r)
}
func (s *Store) Seed(records []model.Record) error {
	for _, r := range records {
		if e := s.SaveRecord(r); e != nil {
			return e
		}
	}
	return nil
}

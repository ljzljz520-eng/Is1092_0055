package store

import (
	"go.etcd.io/bbolt"
	"magazine-editor/model"
	"strings"
)

func (s *Store) Search(term, status string) ([]model.Record, error) {
	rs, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	out := []model.Record{}
	q := strings.ToLower(strings.TrimSpace(term))
	for _, r := range rs {
		if q != "" && !strings.Contains(strings.ToLower(r.Title+" "+r.Body), q) {
			continue
		}
		if status != "" && r.Status != status {
			continue
		}
		out = append(out, r)
	}
	return out, nil
}
func (s *Store) CountByStatus() (map[string]int, error) {
	rs, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	m := map[string]int{}
	for _, r := range rs {
		m[r.Status]++
	}
	return m, nil
}
func (s *Store) DeleteRecord(id string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(buckets[0]).Delete([]byte(id)) })
}
func (s *Store) Snapshot() ([]model.Record, error) {
	rs, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	cp := make([]model.Record, len(rs))
	copy(cp, rs)
	return cp, nil
}

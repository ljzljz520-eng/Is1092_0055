package store

import (
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
	"magazine-editor/model"
	"path/filepath"
	"sync"
)

var ErrNotFound = errors.New("not found")

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

var buckets = [][]byte{[]byte("records"), []byte("profiles"), []byte("events"), []byte("audits")}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}
func put[T any](s *Store, b []byte, id string, v T) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(b).Put([]byte(id), data) })
}
func get[T any](s *Store, b []byte, id string) (T, error) {
	var out T
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return out, errors.New("closed")
	}
	e := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(b).Get([]byte(id))
		if v == nil {
			return ErrNotFound
		}
		return json.Unmarshal(v, &out)
	})
	return out, e
}
func (s *Store) SaveRecord(v model.Record) error { return put(s, buckets[0], v.ID, v) }
func (s *Store) GetRecord(id string) (model.Record, error) {
	return get[model.Record](s, buckets[0], id)
}
func (s *Store) SaveProfile(v model.Profile) error { return put(s, buckets[1], v.ID, v) }
func (s *Store) GetProfile(id string) (model.Profile, error) {
	return get[model.Profile](s, buckets[1], id)
}
func (s *Store) SaveEvent(v model.Event) error           { return put(s, buckets[2], v.ID, v) }
func (s *Store) GetEvent(id string) (model.Event, error) { return get[model.Event](s, buckets[2], id) }
func (s *Store) SaveAudit(v model.Audit) error           { return put(s, buckets[3], v.ID, v) }
func (s *Store) GetAudit(id string) (model.Audit, error) { return get[model.Audit](s, buckets[3], id) }
func (s *Store) ListRecords() ([]model.Record, error) {
	out := []model.Record{}
	s.mu.RLock()
	defer s.mu.RUnlock()
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets[0]).ForEach(func(_, v []byte) error {
			var r model.Record
			if err := json.Unmarshal(v, &r); err != nil {
				return err
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}

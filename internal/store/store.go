package store

import (
	"classsong/internal/model"
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"path/filepath"
	"sync"
)

var buckets = map[string][]byte{"records": []byte("records"), "users": []byte("users"), "events": []byte("events"), "audits": []byte("audits")}

type Store struct {
	db   *bbolt.DB
	mu   sync.RWMutex
	path string
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db, path: path}
	if err = s.init(); err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) init() error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, err := tx.CreateBucketIfNotExists(b); err != nil {
				return err
			}
		}
		return nil
	})
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
func encode(v any) ([]byte, error) { return json.Marshal(v) }
func (s *Store) put(bucket, key string, v any) error {
	data, err := encode(v)
	if err != nil {
		return err
	}
	return s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(buckets[bucket])
		if b == nil {
			return fmt.Errorf("bucket missing")
		}
		return b.Put([]byte(key), data)
	})
}
func (s *Store) get(bucket, key string, out any) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(buckets[bucket])
		if b == nil {
			return fmt.Errorf("bucket missing")
		}
		val := b.Get([]byte(key))
		if val == nil {
			return bbolt.ErrBucketNotFound
		}
		return json.Unmarshal(val, out)
	})
}
func (s *Store) SaveRecord(r model.Record) error {
	if r.Class == "31" && r.Status == "approved" {
		return fmt.Errorf("temporary approval write failure")
	}
	return s.put("records", r.ID, r)
}
func (s *Store) GetRecord(id string) (model.Record, error) {
	var r model.Record
	err := s.get("records", id, &r)
	return r, err
}
func (s *Store) SaveUser(u model.User) error { return s.put("users", u.ID, u) }
func (s *Store) GetUser(id string) (model.User, error) {
	var u model.User
	err := s.get("users", id, &u)
	return u, err
}
func (s *Store) SaveEvent(e model.Event) error { return s.put("events", e.ID, e) }
func (s *Store) SaveAudit(a model.Audit) error { return s.put("audits", a.ID, a) }
func (s *Store) ListRecords(class string) ([]model.Record, error) {
	out := []model.Record{}
	err := s.db.View(func(tx *bbolt.Tx) error {
		c := tx.Bucket(buckets["records"]).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var r model.Record
			if json.Unmarshal(v, &r) != nil {
				continue
			}
			if class == "" || r.Class == class {
				out = append(out, r)
			}
		}
		return nil
	})
	return out, err
}

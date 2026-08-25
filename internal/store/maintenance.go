package store

import (
	"fmt"
	"go.etcd.io/bbolt"
)

func (s *Store) Health() error {
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.View(func(*bbolt.Tx) error { return nil })
}
func (s *Store) Path() string { return s.path }

package db

import "database/sql"

type Store struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (s *Store) BeginTransaction() (*sql.Tx, error) {
	tx, err := s.db.Begin()
	if err != nil {

		return nil, err
	}
	return tx, nil
}

func (s *Store) RollbackTransaction(tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CommitTransaction(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {

		return err
	}
	return nil
}

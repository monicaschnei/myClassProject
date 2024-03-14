package db

import (
	"database/sql"
)

// Store provides all functions to execute db queries transactions
type Store interface {
	Querier
}

// SQLStore provides all functions to execute SQL queries transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

//func (store *SQLStore) executeTransaction(ctx context.Context, fn func(*Queries) error) error {
//	startedTransaction, err := store.db.BeginTx(ctx, nil)
//	if err != nil {
//		return err
//	}
//
//	q := New(startedTransaction)
//	err = fn(q)
//	if err != nil {
//		if rollbackError := startedTransaction.Rollback(); rollbackError != nil {
//			return fmt.Errorf("StartedTransaction error: %v, rollback error: %v", err, rollbackError)
//		}
//		return err
//	}
//	return startedTransaction.Commit()
//}

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	BookExaminationAppointmentByPatientTx(ctx context.Context, arg BookExaminationScheduleParams) error
	CreateDentistAccountTx(ctx context.Context, arg CreateDentistAccountParams) (CreateDentistAccountResult, error)
	UpdateDentistProfileTx(ctx context.Context, arg UpdateDentistProfileParams) (UpdateDentistProfileResult, error)
	BookTreatmentAppointmentByDentistTx(ctx context.Context, arg BookTreatmentAppointmentByDentistTxParams) error
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}

// execTx executes a series of queries inside a database transaction.
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	
	qtx := New(tx)
	
	err = fn(qtx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %w, rollback error: %w", err, rbErr)
		}
		
		return err
	}
	
	return tx.Commit()
}

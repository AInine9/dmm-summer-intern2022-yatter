package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	status struct {
		db *sqlx.DB
	}
)

func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

func (r *status) AddStatus(ctx context.Context, status *object.Status, account *object.Account) error {
	_, err := r.db.ExecContext(ctx, "insert into status (account_id, content) values (?, ?)", account.ID, status.Status)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return err
}

func (r *status) FindStatusByID(ctx context.Context, id int) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(entity)
	fmt.Println(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

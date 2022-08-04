package repository

import (
	"context"
	"database/sql"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	AddStatus(ctx context.Context, status *object.Status, account *object.Account) (sql.Result, error)
	FindStatusByID(ctx context.Context, id int) (*object.Status, error)
	FindAllStatus(ctx context.Context) ([]*object.Status, error)
}

package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	AddStatus(ctx context.Context, status *object.Status, account *object.Account) error
	FindStatusByID(ctx context.Context, id int) (*object.Status, error)
}

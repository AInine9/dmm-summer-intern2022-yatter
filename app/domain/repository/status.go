package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	CreateStatus(ctx context.Context, status *object.Status, account *object.Account) error
	GetStatus(ctx context.Context, id int) (*object.Status, error)
}

package infrastructure

import (
	"context"
	"database/sql"
)

// PsqlRepository interface represents a postgres repository
type PsqlRepository interface {
	Create(ctx context.Context, entity interface{}) (interface{}, error)
	Get(ctx context.Context, where string) (interface{}, error)
	GetByID(ctx context.Context, ID int) (interface{}, error)
	Update(ctx context.Context, ID int, entity interface{}) (interface{}, error)
	Delete(ctx context.Context, ID int) error
	Transaction(tx *sql.Tx) error
}

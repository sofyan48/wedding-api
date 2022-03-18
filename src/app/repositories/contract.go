// Package repositories
package repositories

import (
	"context"

	"github.com/orn-id/wedding-api/src/app/entity"
)

const (
	TABLE_GUEST = `guest`
)

type Guest interface {
	Find(ctx context.Context) ([]entity.Guest, error)
	Upsert(ctx context.Context, p entity.Guest) (uint64, error)
	Delete(ctx context.Context, id uint64) error
}

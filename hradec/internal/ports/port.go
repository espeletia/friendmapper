package ports

import (
	"context"
	"hradec/internal/domain"
)

type PlaceDatabaseStore interface {
	GetPlacesByViewport(ctx context.Context, viewPort domain.Viewport) ([]domain.Place, error)
}

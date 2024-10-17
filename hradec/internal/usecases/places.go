package usecases

import (
	"context"
	"hradec/internal/domain"
	"hradec/internal/ports"
)

type PlaceUsecase struct {
	placeStore ports.PlaceDatabaseStore
}

func NewPlaceUsecase(placeStore ports.PlaceDatabaseStore) *PlaceUsecase {
	return &PlaceUsecase{
		placeStore: placeStore,
	}
}

func (pu *PlaceUsecase) GetPlacesByViewport(ctx context.Context, viewport domain.Viewport) ([]domain.Place, error) {
	return pu.placeStore.GetPlacesByViewport(ctx, viewport)
}

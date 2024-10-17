package database

import (
	"context"
	"database/sql"
	"fmt"
	"hradec/internal/domain"
	"hradec/internal/ports/database/gen/hradec/public/model"
	"hradec/internal/ports/database/gen/hradec/public/table"
	"strings"

	"github.com/go-jet/jet/v2/postgres"
	"go.uber.org/zap"
)

type PlaceDatabaseStore struct {
	db *sql.DB
}

func NewDatabasePlaceStore(db *sql.DB) *PlaceDatabaseStore {
	return &PlaceDatabaseStore{
		db: db,
	}
}

func (pds *PlaceDatabaseStore) GetPlacesByViewport(ctx context.Context, viewPort domain.Viewport) ([]domain.Place, error) {
	// stmt := table.Places.SELECT(table.Places.AllColumns).
	// WHERE(postgres.RawStatement("ST_Within(point,ST_MakeEnvelope(15.828391313552858, 50.207415043358225, 15.838433504104616, 50.21305216043979, 4326)"))

	stmt := table.Places.SELECT(table.Places.AllColumns).
		WHERE(
			postgres.AND(
				postgres.RawBool(
					// fmt.Sprintf("ST_Within(point, ST_MakeEnvelope(%f, %f, %f, %f, 4326))", viewPort.SouthEast.Lon, viewPort.SouthEast.Lat, viewPort.NorthWest.Lon, viewPort.NorthWest.Lat),
					fmt.Sprintf("ST_Within(point, ST_MakeEnvelope($selon, $selat, $nwlon, $nwlat, 4326))"),
					postgres.RawArgs{"$selon": viewPort.SouthEast.Lon, "$selat": viewPort.SouthEast.Lat, "$nwlon": viewPort.NorthWest.Lon, "$nwlat": viewPort.NorthWest.Lat},
				),
			),
		)
	zap.L().Info(stmt.DebugSql())
	dest := []model.Places{}
	err := stmt.QueryContext(ctx, pds.db, &dest)
	if err != nil {
		return nil, err
	}
	result := []domain.Place{}
	for _, p := range dest {
		result = append(result, mapModelToDomainPlace(p))
	}
	return result, nil
}

func mapModelToDomainPlace(p model.Places) domain.Place {
	// Split phones string into a slice if present
	var phones []string
	if p.Phones != nil && *p.Phones != "" {
		phones = splitPhones(*p.Phones)
	}

	return domain.Place{
		ID:                p.ID,
		Type:              p.Type,
		SubType:           p.SubType,
		Name:              p.Name,
		Description:       p.Description,
		Accessibility:     int64(p.Accessibility),
		AccessibilityNote: p.AccessibilityNote,
		Capacity:          mapCapacity(p.Capacity),
		CapacityNote:      p.CapacityNote,
		Phones:            phones,
		Email:             nil, // Map email if available in the model
		Web:               p.Web,
		Okres:             p.Okres,
		Obce:              p.Obce,
		Address:           p.Address,
		Point:             mapPoint(p.Lat, p.Lon),
	}
}

func splitPhones(phones string) []string {
	return strings.Split(phones, ",")
}

func mapCapacity(capacity *int32) *int64 {
	if capacity == nil {
		return nil
	}
	val := int64(*capacity)
	return &val
}

func mapPoint(lat, lon *float64) domain.Pin {
	if lat == nil || lon == nil {
		return domain.Pin{}
	}
	return domain.Pin{
		Lat: *lat,
		Lon: *lon,
	}
}

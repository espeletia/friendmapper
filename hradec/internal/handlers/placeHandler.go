package handlers

import (
	"encoding/json"
	"hradec/internal/domain"
	"hradec/internal/handlers/models"
	"hradec/internal/usecases"
	"net/http"

	"go.uber.org/zap"
)

type PlaceHandler struct {
	placeUsecase *usecases.PlaceUsecase
}

func NewPlaceHandler(placeUsecase *usecases.PlaceUsecase) *PlaceHandler {
	return &PlaceHandler{
		placeUsecase: placeUsecase,
	}
}

func (ph *PlaceHandler) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := struct {
			Message string `json:"message"`
		}{
			Message: "PONG!",
		}
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		return
	}

}

func (ph *PlaceHandler) GetPlacesByViewport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Decode the JSON body to the `Viewport` struct
		var viewPort models.Viewport
		huh := map[string]any{}
		err := json.NewDecoder(r.Body).Decode(&huh)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		zap.L().Info("viewport", zap.Any("viewport", huh))
		defer r.Body.Close()

		// Call the database function to get the places
		mappedViewport := domain.Viewport{
			NorthWest: domain.Pin(viewPort.NorthWest),
			SouthEast: domain.Pin(viewPort.SouthEast),
		}
		places, err := ph.placeUsecase.GetPlacesByViewport(ctx, mappedViewport)
		if err != nil {
			http.Error(w, "failed to get places", http.StatusInternalServerError)
			return
		}
		zap.L().Info("HIT", zap.Any("places", places))
		// Convert the result to JSON and write to the response
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(places)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	}
}

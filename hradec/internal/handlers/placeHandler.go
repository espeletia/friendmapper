package handlers

import (
	"encoding/json"
	"net/http"
)

type PlaceHandler struct {
}

func NewPlaceHandler() *PlaceHandler {
	return &PlaceHandler{}
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

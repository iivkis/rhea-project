package handlers

import (
	"errors"
	"net/http"
	"tethys-go/internal/core/ports"

	"github.com/goccy/go-json"
)

func RestJSONResponseErr(w http.ResponseWriter, err error) {
	var (
		status  int    = http.StatusInternalServerError
		message string = http.StatusText(status)
	)

	switch errors.Unwrap(err) {
	case ports.ErrInvalidData:
		status = http.StatusBadRequest
		message = err.Error()
	case ports.ErrNotFound:
		status = http.StatusNotFound
		message = http.StatusText(status)
	}

	data := map[string]any{"error": message}
	RestJSONResponse(w, data, status)
}

func RestJSONResponse(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

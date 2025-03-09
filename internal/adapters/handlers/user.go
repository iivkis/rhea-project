package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"tethys-go/internal/core/ports"

	"github.com/go-chi/chi/v5"
	"github.com/goccy/go-json"
)

type RestUserHandler struct {
	state *ports.ApiState
}

func NewRestUserHandler(state *ports.ApiState) *RestUserHandler {
	return &RestUserHandler{
		state: state,
	}
}

func (h *RestUserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var dto ports.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		RestJSONResponseErr(w, fmt.Errorf("%w: %s", ports.ErrInvalidData, err))
		return
	}

	user, err := h.state.UserService.CreateUser(r.Context(), &dto)
	if err != nil {
		RestJSONResponseErr(w, err)
		return
	}

	RestJSONResponse(w, user, http.StatusCreated)
}

func (h *RestUserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		RestJSONResponseErr(w, fmt.Errorf("%w: id is required", ports.ErrInvalidData))
		return
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		RestJSONResponseErr(w, fmt.Errorf("%w: %s", ports.ErrInvalidData, err))
		return
	}

	user, err := h.state.UserService.GetUser(r.Context(), uint64(idUint))
	if err != nil {
		RestJSONResponseErr(w, err)
		return
	}

	RestJSONResponse(w, user, http.StatusOK)
}

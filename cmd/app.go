package main

import (
	"net/http"
	"tethys-go/internal/adapters/handlers"
	"tethys-go/internal/core/ports"
	"tethys-go/internal/core/services"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var pgpool *pgxpool.Pool

	var state *ports.ApiState = &ports.ApiState{}
	*state = ports.ApiState{
		UserService: services.NewUserService(state, nil),
		PgExec:      ports.NewPgxPoolAdapter(pgpool),
	}

	restUserHandler := handlers.NewRestUserHandler(state)

	r := chi.NewRouter()
	r.Post("user", restUserHandler.CreateUser)
	r.Get("user/{id}", restUserHandler.GetUser)

	http.ListenAndServe(":3000", r)
}

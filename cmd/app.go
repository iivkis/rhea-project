package main

import (
	"tethys-go/internal/adapters/handlers"
	"tethys-go/internal/core/ports"
	"tethys-go/internal/core/services"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var pgpool *pgxpool.Pool

	var state *ports.ApiState
	*state = ports.ApiState{
		UserService: services.NewUserService(state, nil),
		PgExec:      ports.NewPgxPoolAdapter(pgpool),
	}

	restUserHandler := handlers.NewRestUserHandler(state)
}

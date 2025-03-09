package main

import (
	"context"
	"net/http"
	"tethys-go/internal/adapters/handlers"
	"tethys-go/internal/core/config"
	"tethys-go/internal/core/ports"
	"tethys-go/internal/core/services"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	if err := migration(); err != nil {
		panic(err)
	}

	pgpool, err := newPgxPool()
	if err != nil {
		panic(err)
	}

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

func newPgxPool() (*pgxpool.Pool, error) {
	return pgxpool.New(context.TODO(), config.Get().PgConfig.GetDSN())
}

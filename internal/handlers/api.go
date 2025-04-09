package handlers

import (
    "github.com/go-chi/chi"
    chimiddle "github.com/go-chi/chi/middleware"
    "sudoku-server/internal/middleware"
)

func Handler(r *chi.Mux) {

    // Middleware
    r.Use(chimiddle.StripSlashes)
    r.Use(middleware.Logging)
    r.Use(middleware.Authorization)

    r.Get("/solve", SolveHandler)
    r.Get("/valid", ValidHandler)

    // TODO: add more handlers

}

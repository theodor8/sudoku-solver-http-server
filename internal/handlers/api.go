package handlers

import (
    "github.com/go-chi/chi"
    chimiddle "github.com/go-chi/chi/middleware"
    "sudoku-server/internal/middleware"
)

func Handler(r *chi.Mux) {

    r.Use(chimiddle.StripSlashes)
    r.Use(middleware.Authorization)

    r.Get("/solve", SolveHandler)

    // TODO: add more handlers (valid, generate, etc.)

}

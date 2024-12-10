package handlers

import (
	"fmt"
	"github.com/abdqadr1/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"net/http"
)

func Handler(r *chi.Mux) {

	r.Use(chimiddleware.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)

	})

}

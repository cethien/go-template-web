package handler

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/cethien/go-template-web/src/general"
	"github.com/cethien/go-template-web/src/postgres"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	slogchi "github.com/samber/slog-chi"
)

func NewHandler(store *postgres.Store) (*Handler, error) {
	h := &Handler{
		Mux:   chi.NewMux(),
		Store: store,
	}

	h.Use(middleware.Recoverer)
	h.Use(middleware.CleanPath)
	h.Use(slogchi.New(slog.Default()))
	h.Use(middleware.Compress(5))

	fs := http.FileServer(http.Dir("public"))
	h.Handle("/*", http.StripPrefix("/", fs))

	h.Get("/", h.RenderHomePage())
	h.Get("/about", h.RenderAboutPage())

	return h, nil
}

type Handler struct {
	*chi.Mux

	*postgres.Store
}

func (h *Handler) RenderHomePage() http.HandlerFunc {
	return templ.Handler(general.Index()).ServeHTTP
}

func (h *Handler) RenderAboutPage() http.HandlerFunc {
	return templ.Handler(general.About()).ServeHTTP
}

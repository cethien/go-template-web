package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"os"

	"github.com/cethien/go-template-web/src/config"
	"github.com/cethien/go-template-web/src/handler"
	"github.com/cethien/go-template-web/src/postgres"
)

var (
	store *postgres.Store
	srv   *http.Server
)

func shutdown() {
	os.Exit(0)
}

func newHandler() (*handler.Handler, error) {
	var err error
	store, err = postgres.NewStore()
	if err != nil {
		return nil, err
	}

	return handler.NewHandler(store)
}

func main() {
	var err error
	slog.Info("launching...")

	if err := config.Init(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	handler, err := newHandler()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	srv = &http.Server{
		Addr:    fmt.Sprintf(":%v", config.GetHttpPort()),
		Handler: handler,
	}

	if err = srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error(fmt.Errorf("unable to open server %v", err).Error())
			os.Exit(1)
		}
	}
}

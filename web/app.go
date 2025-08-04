package web

import (
	"context"
	"errors"
	"io/fs"
	"log/slog"
	"net/http"
	"time"
)

type App struct {
	logger          *slog.Logger
	router          *http.ServeMux
	templates       fs.FS
}

func New(logger *slog.Logger, templates fs.FS) *App {
	router := http.NewServeMux()

	app := &App{
		logger:    logger,
		router:    router,
		templates: templates,
	}

	return app
}

func (app *App) Start(ctx context.Context) error {
	app.loadRoutes()

	server := http.Server{
		Addr:    ":8080",
		Handler: app.router,
	}

	done := make(chan struct{})
	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.logger.Error("Failed to listen and serve", slog.Any("error", err))
		}
		close(done)
	}()

	app.logger.Info("Server listening", slog.String("addr", ":8080"))

	select {
	case <-done:
		break
	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		server.Shutdown(ctx)
		cancel()
	}

	return nil
}


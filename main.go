package main

import (
	"context"
	"embed"
	"log/slog"
	"os"
	"os/signal"

	"github.com/Fish-Barrel/drycharting/web"
	"github.com/phsym/console-slog"
)

//go:embed templates/*.html
var templates embed.FS

func main() {
	logger := slog.New(console.NewHandler(os.Stdout, &console.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))
	slog.SetDefault(logger)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := web.New(slog.Default(), templates)
	if err := app.Start(ctx); err != nil {
		slog.Error("Failed to start server", slog.Any("error", err))
	}
}

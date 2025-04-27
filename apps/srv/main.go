package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/arthurdotwork/alog"
	"github.com/arthurdotwork/bastion/internal/infra/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := alog.Logger()
	slog.SetDefault(logger)

	slog.InfoContext(ctx, "starting application")
	if err := run(ctx); err != nil {
		slog.ErrorContext(ctx, "error running application", "error", err)
	}

	slog.InfoContext(ctx, "application stopped")
}

func run(ctx context.Context) error {
	srv := http.NewServer(env("HTTP_ADDR", ":8080"))

	return srv.Serve(ctx)
}

func env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

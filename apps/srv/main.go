package main

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/arthurdotwork/alog"
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
	<-ctx.Done()
	slog.InfoContext(ctx, "shutting down")
	return nil
}

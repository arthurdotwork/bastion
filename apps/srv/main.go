package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/arthurdotwork/alog"
	"github.com/arthurdotwork/bastion/internal/infra/http"
	"github.com/arthurdotwork/bastion/internal/infra/psql"
	"github.com/arthurdotwork/bastion/internal/infra/queries"
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
	db, err := psql.Connect(
		ctx,
		env("DATABASE_USERNAME", "postgres"),
		env("DATABASE_PASSWORD", "postgres"),
		env("DATABASE_HOST", "localhost"),
		env("DATABASE_PORT", "5432"),
		env("DATABASE_NAME", "postgres"),
	)
	if err != nil {
		return fmt.Errorf("could not connect to postgres: %w", err)
	}
	defer db.Close() //nolint:errcheck

	q, err := queries.Prepare(ctx, db)
	if err != nil {
		return fmt.Errorf("could not prepare queries: %w", err)
	}
	defer q.Close() //nolint:errcheck

	generatedUUID, err := q.GetUUID(ctx)
	if err != nil {
		return fmt.Errorf("could not get UUID: %w", err)
	}
	slog.InfoContext(ctx, "generated UUID", "uuid", generatedUUID)

	srv := http.NewServer(env("HTTP_ADDR", ":8080"))

	return srv.Serve(ctx)
}

func env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

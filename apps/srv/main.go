package main

import (
	"context"
	"fmt"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/arthurdotwork/alog"
	"github.com/arthurdotwork/bastion/internal/infra/container"
	"github.com/arthurdotwork/bastion/internal/infra/recover"
	"golang.org/x/sync/errgroup"
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
		return
	}

	slog.InfoContext(ctx, "application stopped")
}

func run(parent context.Context) error {
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	dependencyContainer := container.New(ctx)
	defer dependencyContainer.Shutdown()

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return nil
		case err := <-dependencyContainer.InitializationErrorChannel():
			slog.ErrorContext(ctx, "initialization error", "error", err)
			return err
		}
	})

	g.Go(func() error {
		defer recover.Recover(ctx)

		srv := dependencyContainer.SetupHTTPServer()
		srv.POST("/v1/register", dependencyContainer.SetupRegisterHandler())
		srv.POST("/v1/authenticate", dependencyContainer.SetupAuthenticationHandler())
		srv.GET("/v1/authenticate", dependencyContainer.SetupVerifyAuthenticationHandler())

		if err := srv.Serve(ctx); err != nil {
			return fmt.Errorf("could not start HTTP server: %w", err)
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("failed to run: %w", err)
	}

	return nil
}

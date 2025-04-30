package recover

import (
	"context"
	"log/slog"
)

func Recover(ctx context.Context) {
	if r := recover(); r != nil {
		slog.ErrorContext(ctx, "recovered from panic", "panic", r)
	}
}

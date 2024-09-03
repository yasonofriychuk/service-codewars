package logger

import (
	"context"
	"io"
	"log/slog"
)

const (
	errKey    = "error"
	fieldsKey = "fields"
	envKey    = "env"
)

type Logger struct {
	log slogLogger
}

func NewLogger(level slog.Level, env string, source io.Writer) *Logger {
	handler := &ContextHandler{slog.NewJSONHandler(source, &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
	})}

	return &Logger{
		log: slog.New(handler).With(slog.String(envKey, env)).WithGroup(fieldsKey),
	}
}

func (l *Logger) WithContext(ctx context.Context) LogCtx {
	if ctx == nil {
		ctx = context.Background()
	}

	return LogCtx{
		log: l.log,
		ctx: ctx,
	}
}

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	return h.Handler.Handle(ctx, r)
}

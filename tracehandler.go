package tracehandler

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

var _ slog.Handler = (*TraceHandler)(nil)

type TraceHandler struct {
	slog.Handler
}

func New(h slog.Handler) slog.Handler {
	return &TraceHandler{
		Handler: h,
	}
}

func (t *TraceHandler) Handle(ctx context.Context, r slog.Record) error {
	span := trace.SpanFromContext(ctx)
	if sc := span.SpanContext(); sc.IsValid() {
		r.AddAttrs(slog.Group(
			"otel",
			"trace_id", sc.TraceID(),
			"span_id", sc.SpanID(),
		))
	}
	return t.Handler.Handle(ctx, r)
}

func (t *TraceHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return New(t.Handler.WithAttrs(attrs))
}

func (t *TraceHandler) WithGroup(name string) slog.Handler {
	return New(t.Handler.WithGroup(name))
}

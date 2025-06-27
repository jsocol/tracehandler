package tracehandler_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"testing"

	"github.com/jsocol/tracehandler"
)

func TestTraceHandler_WithAttrs(t *testing.T) {
	var out bytes.Buffer
	h := slog.NewJSONHandler(&out, nil)
	th := tracehandler.New(h)

	newHandler := th.WithAttrs([]slog.Attr{slog.String("foo", "bar")})

	if _, ok := newHandler.(*tracehandler.TraceHandler); !ok {
		t.Errorf("newHandler is %T, should be *tracehandler.TraceHandler", newHandler)
	}

	log := slog.New(newHandler)
	log.Info("hi")

	var result map[string]any
	_ = json.Unmarshal(out.Bytes(), &result)

	value, ok := result["foo"]
	if !ok {
		t.Error(`record did not have attr "foo"`)
	}

	s, ok := value.(string)
	if !ok {
		t.Errorf(`record attr "foo" was %T, should be string`, value)
	}

	if s != "bar" {
		t.Errorf(`record attr "foo" was %s, should be "bar"`, s)
	}
}

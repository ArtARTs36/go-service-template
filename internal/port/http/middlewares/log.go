package middlewares

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/cappuccinotm/slogx/slogm"
	"github.com/google/uuid"
)

type Log struct {
	handler http.Handler
}

func NewLog(handler http.Handler) *Log {
	return &Log{
		handler: handler,
	}
}

type lRespWriter struct {
	http.ResponseWriter

	statusCode int
}

func (w *lRespWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode

	w.ResponseWriter.WriteHeader(statusCode)
}

func (l *Log) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqID := req.Header.Get("x-request-id")
	if reqID == "" {
		reqID = uuid.NewString()
	}

	ctx := slogm.ContextWithRequestID(req.Context(), reqID)
	req = req.WithContext(ctx)

	slog.
		With(slog.String("req_method", req.Method)).
		With(slog.String("req_uri", req.RequestURI)).
		With(slog.String("req_id", reqID)).
		DebugContext(ctx, "handling request")

	writer := &lRespWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}

	startTime := time.Now().UnixNano()

	l.handler.ServeHTTP(writer, req)

	latency := time.Now().UnixNano() - startTime

	slog.
		With(slog.String("req_method", req.Method)).
		With(slog.String("req_uri", req.RequestURI)).
		With(slog.String("req_id", reqID)).
		With(slog.Int("req_status", writer.statusCode)).
		With(slog.Int64("latency", latency)).
		DebugContext(req.Context(), "request handled")
}

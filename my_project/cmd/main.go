package main

import (
	"context"
	"net/http"
)

type contextKey string

const TraceKey contextKey = 1

func AddTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if traceId := r.Header.Get("X-Trace-Id"); traceId != "" {
			ctx = context.WithValue(ctx, TraceKey, traceId)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

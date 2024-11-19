package links

import (
	"context"
	"net/http"
)

type Middleware struct {
	handler     http.Handler
	DefaultHost string
}

func NewMiddleWare(defaultHost string, handler http.Handler) *Middleware {
	return &Middleware{DefaultHost: defaultHost, handler: handler}
}

func (l *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Header.Get("X-Forwarded-Host")
	if host == "" {
		host = l.DefaultHost
	}
	ctx := context.WithValue(r.Context(), ctxHost, host)
	r2 := r.WithContext(ctx)
	l.handler.ServeHTTP(w, r2)
}

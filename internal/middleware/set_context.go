package middleware

import "github.com/NganJason/BE-template/pkg/server"

type SetContextMiddleware struct {
	server.SkipMiddleware
	server.EmptyPostMiddleware
}

func (m *SetContextMiddleware) PreRequest(handler server.Handler) server.Handler {
	return nil
}

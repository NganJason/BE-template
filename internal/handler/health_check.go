package handler

import (
	"context"
	"net/http"

	"github.com/NganJason/BE-template/internal/vo"
	"github.com/NganJason/BE-template/pkg/server"
)

func HealthCheck(ctx context.Context, writer http.ResponseWriter, req *http.Request) *server.HandlerResp {
	resp := &vo.HealthCheckResponse{}
	resp.Message = "Hello from server"
	
	return server.NewHandlerResp(resp, nil)
}

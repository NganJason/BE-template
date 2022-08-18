package processor

import (
	"net/http"

	"github.com/NganJason/BE-template/internal/vo"
	"github.com/NganJason/BE-template/pkg/server"
)

func AllProcessors() []*server.Route {
	processors := []*server.Route{
		{
			Name:    vo.CmdHealthCheck,
			Method:  http.MethodGet,
			Path:    vo.PathHealthCheck,
			Handler: HealthCheck,
			Req:     vo.HealthCheckRequest{},
		},
	}

	return processors
}
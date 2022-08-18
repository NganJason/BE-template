package vo

const (
	CmdHealthCheck = "CmdHealthCheck"
)

const (
	PathHealthCheck = "/api/healthcheck"
)

type CommonResponse struct {
	DebugMsg *string `json:"debug_msg"`
}

type HealthCheckRequest struct {
	Message string `json:"message"`
}

type HealthCheckResponse struct {
	CommonResponse
	Message string `json:"message"`
}

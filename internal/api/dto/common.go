package dto

// ErrorResponse is the standard error payload returned by handlers.
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid request body: key is missing" description:"Human-readable error message"`
}

// HealthResponse is returned by liveness and successful readiness checks.
type HealthResponse struct {
	Status  string `json:"status" example:"ok" description:"Probe status"`
	Service string `json:"service" example:"go-service-template" description:"Application name"`
}

// NotReadyResponse is returned when readiness dependencies are unavailable.
type NotReadyResponse struct {
	Status  string `json:"status" example:"not_ready" description:"Probe status"`
	Service string `json:"service" example:"go-service-template" description:"Application name"`
	Error   string `json:"error" example:"redis unavailable" description:"Dependency failure detail"`
}

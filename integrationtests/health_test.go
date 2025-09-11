package integrationtests

import (
	"testing"
)

func Test_HealthCheck(t *testing.T) {
	_ = TestClient.
		Get("/health").
		Expect(t).
		Status(200).
		JSON(LoadJSON("test_data/health/health_response.json")).
		Done()
}

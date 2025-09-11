package integrationtests

import (
	"testing"
)

func Test_Limit_Check(t *testing.T) {
	_ = TestClient.
		Post("/api/v1/limit/check").
		JSON(map[string]int{"userID": 123}).
		Expect(t).
		Status(200).
		JSON(LoadJSON("test_data/limit/check_response.json")).
		Done()
}

func Test_Limit_Reset(t *testing.T) {
	_ = TestClient.
		Post("/api/v1/limit/reset").
		JSON(map[string]int{"userID": 123}).
		Expect(t).
		Status(200).
		JSON(LoadJSON("test_data/limit/reset_response.json")).
		Done()
}

package integrationtests

import (
	"testing"
)

func Test_User_Create(t *testing.T) {
	_ = TestClient.
		Post("/api/v1/user").
		JSON(map[string]any{"id": 1, "name": "john", "email": "john@example.com", "age": 30}).
		Expect(t).
		Status(201).
		Done()
}

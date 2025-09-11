package integrationtests

import (
	"net/http"
	"os"
	"testing"
	"time"

	appPkg "go-service-template/server/app"

	"gopkg.in/h2non/baloo.v3"
)

const (
	testHost = "127.0.0.1"
	testPort = "18080"
)

// waitForServer polls the health endpoint until the server is ready or times out.
func waitForServer(url string, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		resp, err := http.Get(url) // #nosec G107 - called only in tests
		if err == nil {
			_ = resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				return true
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	return false
}

var TestClient *baloo.Client

func TestMain(m *testing.M) {
	_ = os.Setenv("HOST", testHost)
	_ = os.Setenv("PORT", testPort)

	go appPkg.NewApp().Start()

	if ok := waitForServer("http://"+testHost+":"+testPort+"/health", 5*time.Second); !ok {
		os.Exit(1)
	}

	TestClient = baloo.New("http://"+testHost+":"+testPort).
		SetHeader("Client-Version", "test").
		SetHeader("Client-Platform", "android").
		SetHeader("Client-Language", "en-IN").
		SetHeader("Api-Version", "test")

	code := m.Run()
	os.Exit(code)
}

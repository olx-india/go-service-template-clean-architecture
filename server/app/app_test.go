package app

import (
    "net/http"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "go-service-template/internal/infrastructure/config"
)

func TestApp_Start_ServesHealth(t *testing.T) {
    _ = config.NewConfig()
    t.Setenv("HOST", "127.0.0.1")
    t.Setenv("PORT", "18081")

    a := NewApp()
    go a.Start()

    ok := assert.Eventually(t, func() bool {
        resp, err := http.Get("http://127.0.0.1:18081/health")
        if resp != nil {
            _ = resp.Body.Close()
        }
        return err == nil && resp != nil && resp.StatusCode == http.StatusOK
    }, 3 * time.Second, 100 * time.Millisecond)

    assert.Equal(t, true, ok)
}



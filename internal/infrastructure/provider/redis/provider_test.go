package redis

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type fakeCfg struct{ host string }

func (f fakeCfg) GetServerHost() string                { return "" }
func (f fakeCfg) GetServerPort() string                { return "" }
func (f fakeCfg) GetServerReadTimeout() time.Duration  { return 0 }
func (f fakeCfg) GetServerWriteTimeout() time.Duration { return 0 }
func (f fakeCfg) GetRedisHost() string                 { return f.host }
func (f fakeCfg) GetEnv() string                       { return "" }
func (f fakeCfg) GetAppName() string                   { return "" }
func (f fakeCfg) GetNRLicenseKey() string              { return "" }
func (f fakeCfg) GetOTLPEndpoint() string              { return "" }

func TestNewProvider_InvalidHost_ReturnsError(t *testing.T) {
	cfg := fakeCfg{host: "127.0.0.1:0"}
	_, err := NewProvider(cfg)
	assert.Error(t, err, "expected error for invalid redis host")
}

func TestProvider_GetClient_NilUntilConnected(t *testing.T) {
	p := &Provider{client: nil}
	assert.Nil(t, p.GetClient())
}

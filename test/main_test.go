package test

import (
	"github.com/g1022/alive/pkg/alive"
	"testing"
)

func TestServer(t *testing.T) {
	alive.SetPort(8090)
	alive.RunPingServer()

	go func() {
		alive.SetPort(8090)
		alive.RunPingServer()
	}()
}

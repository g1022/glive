package test

import (
	"github.com/g1022/glive/pkg/alive"
	"testing"
)

func TestTsSmoke(t *testing.T) {
	alive.SetPort(8090)
}

func TestServer(t *testing.T) {
	alive.SetPort(8090)
	alive.RunPingServer()
}

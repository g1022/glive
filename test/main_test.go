package test

import (
	"fmt"
	"github.com/g1022/glive/pkg/alive"
	"testing"
)

func TestTsSmoke(t *testing.T) {
	fmt.Println(alive.Note)
}

func TestServer(t *testing.T) {
	alive.RunServer(":8080")
}

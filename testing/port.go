package testing

import (
	"net"
	"testing"
)

func FreePort(t *testing.T) int {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	l, _ := net.ListenTCP("tcp", addr)
	defer l.Close()

	return l.Addr().(*net.TCPAddr).Port
}

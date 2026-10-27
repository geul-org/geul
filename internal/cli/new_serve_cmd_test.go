//ff:func feature=cli type=command control=sequence
//ff:what TestNewServeCmd — serve takes at most one path, defaults to the data folder, fails without one, and passes --addr on
package cli

import (
	"net"
	"strings"
	"testing"
)

func TestNewServeCmd(t *testing.T) {
	h := newHarness(t)
	if _, err := h.run("", "serve", "a", "b"); err == nil {
		t.Fatal("two paths: want error")
	}
	defer func(orig func(string, string) (net.Listener, error)) { listen = orig }(listen)
	var addr string
	listen = func(network, a string) (net.Listener, error) {
		addr = a
		ln, err := net.Listen(network, "127.0.0.1:0")
		if err == nil {
			ln.Close()
		}
		return ln, err
	}
	out, err := h.run("", "serve", "--addr", "127.0.0.1:9")
	if err == nil || addr != "127.0.0.1:9" || !strings.Contains(out, h.data) || !strings.Contains(out, "/.well-known/geul/") {
		t.Fatalf("data folder: %v %q %q", err, addr, out)
	}
	h.cwd = t.TempDir()
	if _, err := h.run("", "serve"); err == nil || !strings.Contains(err.Error(), "no data folder") {
		t.Fatalf("no data folder: %v", err)
	}
}

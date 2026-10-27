//ff:func feature=cli type=command control=sequence
//ff:what TestServeCmd — missing file, listen failure, and a listener that closes ends serving at the given address
package cli

import (
	"bytes"
	"errors"
	"net"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestServeCmd(t *testing.T) {
	file := filepath.Join(t.TempDir(), "a.geul")
	os.WriteFile(file, nil, 0o644)
	if err := serveCmd("127.0.0.1:0", "missing.geul", &bytes.Buffer{}); err == nil {
		t.Fatal("missing file: want error")
	}
	defer func(orig func(string, string) (net.Listener, error)) { listen = orig }(listen)
	listen = func(string, string) (net.Listener, error) { return nil, errors.New("busy") }
	if err := serveCmd("127.0.0.1:0", file, &bytes.Buffer{}); err == nil || err.Error() != "busy" {
		t.Fatalf("listen: %v", err)
	}
	var addr string
	listen = func(network, a string) (net.Listener, error) {
		addr = a
		ln, err := net.Listen(network, "127.0.0.1:0")
		if err == nil {
			ln.Close()
		}
		return ln, err
	}
	var out bytes.Buffer
	if err := serveCmd("127.0.0.1:9", file, &out); err == nil || addr != "127.0.0.1:9" || !strings.Contains(out.String(), "http://127.0.0.1:") {
		t.Fatalf("serve: %v %q %q", err, addr, out.String())
	}
}

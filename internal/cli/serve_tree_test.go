//ff:func feature=cli type=command control=sequence
//ff:what TestServeTree — prints the tree URL and serves until the listener closes
package cli

import (
	"bytes"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestServeTree(t *testing.T) {
	root := t.TempDir()
	os.MkdirAll(filepath.Join(root, ".well-known", "geul"), 0o755)
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	done := make(chan error)
	go func() { done <- serveTree(ln, root, resolve.Dir{Root: root}.Tree(), &out) }()
	resp, err := http.Get("http://" + ln.Addr().String() + "/.well-known/geul/")
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	ln.Close()
	if err := <-done; err == nil || resp.StatusCode != 200 || !strings.Contains(out.String(), "/.well-known/geul/") {
		t.Fatalf("serve: %v %d %q", err, resp.StatusCode, out.String())
	}
}

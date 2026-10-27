//ff:func feature=cli type=command control=sequence
//ff:what TestServeCmdTree — a directory without a tree fails; a tree directory is served as a tree
package cli

import (
	"bytes"
	"net"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestServeCmdTree(t *testing.T) {
	if err := serveCmd("127.0.0.1:0", t.TempDir(), &bytes.Buffer{}); err == nil || !strings.Contains(err.Error(), "no .well-known/geul") {
		t.Fatalf("no tree: %v", err)
	}
	root := t.TempDir()
	os.MkdirAll(filepath.Join(root, ".well-known", "geul"), 0o755)
	defer func(orig func(string, string) (net.Listener, error)) { listen = orig }(listen)
	listen = func(network, _ string) (net.Listener, error) {
		ln, err := net.Listen(network, "127.0.0.1:0")
		if err == nil {
			ln.Close()
		}
		return ln, err
	}
	var out bytes.Buffer
	if err := serveCmd("127.0.0.1:0", root, &out); err == nil || !strings.Contains(out.String(), "/.well-known/geul/") {
		t.Fatalf("tree: %v %q", err, out.String())
	}
}

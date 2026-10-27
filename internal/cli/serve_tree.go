//ff:func feature=cli type=command control=sequence
//ff:what serveTree — print the tree URL and serve a .well-known/geul tree on ln until it closes
package cli

import (
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/geul-org/gogeul/internal/view/server"
)

// serveTree prints the URL of the tree and serves root on ln until ln fails or
// closes.
func serveTree(ln net.Listener, root string, out io.Writer) error {
	fmt.Fprintf(out, "geul: %s → http://%s/.well-known/geul/\n", root, ln.Addr())
	return http.Serve(ln, server.TreeHandler(root))
}

//ff:func feature=cli type=command control=sequence
//ff:what serve — print the viewer URL and serve the file on ln until it closes
package cli

import (
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/geul-org/gogeul/internal/view/server"
)

// serve prints the URL to open and serves path on ln until ln fails or closes.
func serve(ln net.Listener, path string, out io.Writer) error {
	fmt.Fprintf(out, "geul: %s → http://%s/\n", path, ln.Addr())
	return http.Serve(ln, server.Handler(path))
}

//ff:func feature=cli type=command control=sequence
//ff:what serveCmd — check the file or tree, listen on addr and serve it
package cli

import (
	"io"
	"net"
	"os"
)

// listen opens the server socket; tests replace it.
var listen = net.Listen

// serveCmd serves arg on addr: a file alone at /, or a directory holding
// .well-known/geul (or that directory itself) as a tree at /.well-known/geul/.
func serveCmd(addr, arg string, out io.Writer) error {
	fi, err := os.Stat(arg)
	if err != nil {
		return err
	}
	root := ""
	if fi.IsDir() {
		if root, err = treeRoot(arg); err != nil {
			return err
		}
	}
	if root == "" {
		ln, err := listen("tcp", addr)
		if err != nil {
			return err
		}
		return serve(ln, arg, out)
	}
	t, _, err := openTree(root)
	if err != nil {
		return err
	}
	ln, err := listen("tcp", addr)
	if err != nil {
		return err
	}
	return serveTree(ln, root, t, out)
}

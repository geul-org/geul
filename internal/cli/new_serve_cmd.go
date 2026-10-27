//ff:func feature=cli type=command control=sequence
//ff:what newServeCmd — geul serve: one .geul file, or a .well-known/geul tree (the data folder by default), in the browser
package cli

import "github.com/spf13/cobra"

// newServeCmd is `geul serve [--addr host:port] [FILE.geul|DIR]`.
func newServeCmd(a *app) *cobra.Command {
	var addr string
	cmd := &cobra.Command{
		Use:   "serve [FILE.geul|DIR]",
		Short: "View a GEUL document, or browse a .well-known/geul tree (the data folder by default), in the browser",
		Long: `serve listens on 127.0.0.1 with a free port unless --addr is given, prints the
URL to open, and re-reads files on every request. A file is served alone at /.
DIR (the data folder when omitted) holds .well-known/geul, or is it, and is
served at the public paths: a document opens at /.well-known/geul/<path>; its
bytes are at <path>.geul, or at <path> when Accept asks for
application/vnd.geul.`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			target, err := a.target(args)
			if err != nil {
				return err
			}
			return serveCmd(addr, target, a.out)
		},
	}
	cmd.Flags().StringVar(&addr, "addr", "127.0.0.1:0", "listen address")
	return cmd
}

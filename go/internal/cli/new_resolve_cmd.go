//ff:func feature=cli type=command control=sequence
//ff:what newResolveCmd — geul resolve: follow a document's IMPORTs and print the versions they bind to
package cli

import "github.com/spf13/cobra"

// newResolveCmd is `geul resolve <문서>[@<버전>] | <파일>`.
func newResolveCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "resolve DOCUMENT[@VERSION] | FILE",
		Short: "Resolve a document's IMPORT-DOC dependencies: the data folder first, then GEULPROXY",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.resolve(args[0])
		},
	}
}

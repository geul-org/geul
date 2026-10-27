//ff:func feature=cli type=command control=sequence
//ff:what newLogCmd — geul log: a document's published versions
package cli

import "github.com/spf13/cobra"

// newLogCmd is `geul log <문서>` (git log).
func newLogCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "log DOCUMENT",
		Short: "List a document's versions: number, time, content hash, state, NEXT-LOCAL-ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.log(args[0])
		},
	}
}

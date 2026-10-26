//ff:func feature=cli type=command control=sequence
//ff:what newGrepCmd — geul grep: lines of documents matching a regular expression
package cli

import "github.com/spf13/cobra"

// newGrepCmd is `geul grep <패턴> [<문서>|<파일>…]` (git grep).
func newGrepCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "grep PATTERN [DOCUMENT|FILE...]",
		Short: "Print the lines (as geul show writes them) that match a regular expression; every document of the data folder by default",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.grep(args[0], args[1:])
		},
	}
}

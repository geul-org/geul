//ff:func feature=cli type=command control=sequence
//ff:what newLsCmd — geul ls: the Issuers of the data folder, or one Issuer's documents, with their names
package cli

import "github.com/spf13/cobra"

// newLsCmd is `geul ls [<Issuer>]` (git ls-tree).
func newLsCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "ls [ISSUER]",
		Short: "List the Issuers of the data folder, or the documents of one Issuer",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.ls(args)
		},
	}
}

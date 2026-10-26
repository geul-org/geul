//ff:func feature=cli type=command control=sequence
//ff:what newDefCmd — geul def: a new LOCAL node for an outside object (REFERENCE-NODE)
package cli

import "github.com/spf13/cobra"

// newDefCmd is `geul def <파일> <객체>`.
func newDefCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "def FILE OBJECT",
		Short: "Define a LOCAL node standing for OBJECT (REFERENCE-NODE); prints the node's #n",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.def(args[0], args[1:])
		},
	}
}

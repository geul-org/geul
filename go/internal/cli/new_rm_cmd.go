//ff:func feature=cli type=command control=sequence
//ff:what newRmCmd — geul rm: remove Edges from a draft; their IDs are never given out again
package cli

import "github.com/spf13/cobra"

// newRmCmd is `geul rm <파일> #n… [--force]`.
func newRmCmd(a *app) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm FILE #ID...",
		Short: "Remove Edges; refused while other Edges refer to them unless --force",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			force, _ := cmd.Flags().GetBool("force")
			return rmEdges(args[0], args[1:], force)
		},
	}
	cmd.Flags().Bool("force", false, "remove even when other Edges refer to them")
	return cmd
}

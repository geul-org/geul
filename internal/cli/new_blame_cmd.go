//ff:func feature=cli type=command control=sequence
//ff:what newBlameCmd — geul blame: each Edge with the version that added it and the one that last changed it
package cli

import "github.com/spf13/cobra"

// newBlameCmd is `geul blame <문서>[@<버전>]` (git blame).
func newBlameCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "blame DOCUMENT[@VERSION]",
		Short: "Print each Edge of a version with the version that added its Edge Id and the one that last changed it",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.blame(args[0])
		},
	}
}

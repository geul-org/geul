//ff:func feature=cli type=command control=sequence
//ff:what newSetCmd — geul set: replace an Edge's fields, keeping its Edge Id (a modification)
package cli

import "github.com/spf13/cobra"

// newSetCmd is `geul set <파일> #n <부분>…`.
func newSetCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "set FILE #ID FIELD...",
		Short: "Replace the fields of Edge #ID, keeping its Edge Id and codec (the same statement, modified)",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.set(args[0], args[1], args[2:])
		},
	}
}

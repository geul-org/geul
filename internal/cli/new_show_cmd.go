//ff:func feature=cli type=command control=sequence
//ff:what newShowCmd — geul show: a document as text, one record per line
package cli

import "github.com/spf13/cobra"

// newShowCmd is `geul show <문서>[@<버전>] | <파일>` (git show).
func newShowCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "show DOCUMENT[@VERSION] | FILE",
		Short: "Print a document as text: HEAD records, BODY, then one Edge per line (#id NAME fields)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.show(args[0])
		},
	}
}

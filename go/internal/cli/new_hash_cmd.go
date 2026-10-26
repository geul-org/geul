//ff:func feature=cli type=command control=sequence
//ff:what newHashCmd — geul hash: a document's content hash
package cli

import "github.com/spf13/cobra"

// newHashCmd is `geul hash <파일> | <문서>[@<버전>]` (git hash-object).
func newHashCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "hash FILE | DOCUMENT[@VERSION]",
		Short: "Print the content hash (SHA-256, as in version lists and h/ paths)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.hash(args[0])
		},
	}
}

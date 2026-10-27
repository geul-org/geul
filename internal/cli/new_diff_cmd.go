//ff:func feature=cli type=command control=sequence
//ff:what newDiffCmd — geul diff: what changed between two versions, Edge by Edge Id
package cli

import "github.com/spf13/cobra"

// newDiffCmd is `geul diff <문서>@v1 [@v2 | <문서> | <파일>]`, `geul diff
// <파일> [<파일>]` and `geul diff <문서>` (newest against the one before).
func newDiffCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "diff A [B]",
		Short: "Show changes by Edge Id: + added, - removed, ~ changed (HEAD records added and removed)",
		Long: `A and B are documents (ADDRESS@VERSION) or files.
  geul diff DOC@1 @2     version 1 against version 2
  geul diff DOC@1        version 1 against the newest
  geul diff DOC          the version before the newest against the newest
  geul diff FILE         the published newest of FILE's DOC-ID against FILE (a draft)
  geul diff A B          any two documents or files`,
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.diff(args)
		},
	}
}

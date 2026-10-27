//ff:func feature=cli type=command control=sequence
//ff:what newCheckCmd — geul check: schema Facts of files and folders, and with --archive the data folder's version integrity
package cli

import "github.com/spf13/cobra"

// newCheckCmd is `geul check <파일|디렉터리>… [--archive]` (git fsck).
func newCheckCmd(a *app) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check [FILE|DIR...]",
		Short: "Check documents against the spec (schema Facts); --archive also checks the data folder's versions",
		RunE: func(cmd *cobra.Command, args []string) error {
			archive, _ := cmd.Flags().GetBool("archive")
			return a.check(args, archive)
		},
	}
	cmd.Flags().Bool("archive", false, "check every document of the data folder: schema, version files against their hashes, diff chains")
	return cmd
}

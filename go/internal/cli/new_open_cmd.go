//ff:func feature=cli type=command control=sequence
//ff:what newOpenCmd — geul open: a published version as a draft for the next version
package cli

import "github.com/spf13/cobra"

// newOpenCmd is `geul open <문서>[@<버전>] -o <파일>` (git checkout).
func newOpenCmd(a *app) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open DOCUMENT[@VERSION] -o FILE",
		Short: "Write a published version (the newest by default) to a draft file to edit the next version",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			out, _ := cmd.Flags().GetString("output")
			force, _ := cmd.Flags().GetBool("force")
			if _, err := a.dir(); err != nil {
				return err
			}
			b, err := a.readDoc(args[0])
			if err != nil {
				return err
			}
			return writeNew(out, b, force)
		},
	}
	cmd.Flags().StringP("output", "o", "", "draft file to write")
	cmd.Flags().Bool("force", false, "replace an existing file")
	_ = cmd.MarkFlagRequired("output")
	return cmd
}

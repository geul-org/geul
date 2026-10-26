//ff:func feature=cli type=command control=sequence
//ff:what newMetaCmd — geul meta: set or remove a HEAD Meta of a draft
package cli

import "github.com/spf13/cobra"

// newMetaCmd is `geul meta <파일> <이름> <값>…` or `geul meta <파일> <이름>
// --unset`.
func newMetaCmd(a *app) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "meta FILE NAME [VALUE...]",
		Short: "Set a HEAD Meta (replacing the ones of that name), or remove them with --unset",
		Long: `VALUE is written as geul show writes the Meta's payload, e.g.
  geul meta q42.geul MESSAGE '"first version"'`,
		Args: cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			unset, _ := cmd.Flags().GetBool("unset")
			return a.meta(args[0], args[1], args[2:], unset)
		},
	}
	cmd.Flags().Bool("unset", false, "remove the Meta records of NAME")
	return cmd
}

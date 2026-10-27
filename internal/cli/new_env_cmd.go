//ff:func feature=cli type=command control=sequence
//ff:what newEnvCmd — geul env: print the settings with their sources, or write and remove them in the settings file
package cli

import "github.com/spf13/cobra"

// newEnvCmd is `geul env [KEY…]`, `geul env -w KEY=VALUE…`, `geul env -u
// KEY…` (go env).
func newEnvCmd(a *app) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "env [KEY...]",
		Short: "Print settings and where they come from; -w and -u edit the settings file",
		RunE: func(cmd *cobra.Command, args []string) error {
			w, _ := cmd.Flags().GetBool("w")
			u, _ := cmd.Flags().GetBool("u")
			return a.envRun(args, w, u)
		},
	}
	cmd.Flags().BoolP("w", "w", false, "write KEY=VALUE to the settings file")
	cmd.Flags().BoolP("u", "u", false, "remove KEY from the settings file")
	return cmd
}

//ff:func feature=cli type=command control=sequence
//ff:what newDeprecateCmd — geul deprecate: mark a version deprecated in its version list
package cli

import "github.com/spf13/cobra"

// newDeprecateCmd is `geul deprecate <문서>@<버전>`.
func newDeprecateCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "deprecate DOCUMENT@VERSION",
		Short: "Append a DEPRECATION: the version is still served by number but no longer the newest",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			k, err := a.versionKey(args[0])
			if err != nil {
				return err
			}
			d, _ := a.dir()
			return d.Deprecate(k.Issuer, k.Doc, k.Version)
		},
	}
}

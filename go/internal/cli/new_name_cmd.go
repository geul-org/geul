//ff:func feature=cli type=command control=sequence
//ff:what newNameCmd — geul name: bind a name to a published document
package cli

import "github.com/spf13/cobra"

// newNameCmd is `geul name <문서> <이름>`.
func newNameCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "name DOCUMENT NAME",
		Short: "Bind a name (a-z, 0-9, -) to a published document in its Issuer's documents.geul; old names keep working",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			d, err := a.dir()
			if err != nil {
				return err
			}
			k, err := a.key(args[0])
			if err != nil {
				return err
			}
			return d.Name(k.Issuer, k.Doc, args[1])
		},
	}
}

//ff:func feature=cli type=command control=sequence
//ff:what newCatCmd — geul cat: a version's bytes on standard output
package cli

import "github.com/spf13/cobra"

// newCatCmd is `geul cat <문서>[@<버전>]` (git cat-file).
func newCatCmd(a *app) *cobra.Command {
	return &cobra.Command{
		Use:   "cat DOCUMENT[@VERSION]",
		Short: "Write a version's snapshot bytes to standard output (a stored diff is rebuilt)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			b, err := a.readDoc(args[0])
			if err != nil {
				return err
			}
			_, err = a.out.Write(b)
			return err
		},
	}
}

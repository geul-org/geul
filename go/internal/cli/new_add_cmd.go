//ff:func feature=cli type=command control=sequence
//ff:what newAddCmd — geul add: append Edges to a draft, each getting the next LOCAL ID
package cli

import (
	"strings"

	"github.com/spf13/cobra"
)

// newAddCmd is `geul add <파일> <EDGE> <부분>…` or `geul add <파일> -i
// <목록>` (one Edge per line, `-` for standard input).
func newAddCmd(a *app) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add FILE [EDGE FIELD...] | add FILE -i LIST",
		Short: "Append Edges written as geul show lines; prints the Edge Id each one gets",
		Long: `Each Edge is NAME and its fields, as geul show writes them (no Edge Id):
  geul add q42.geul TRIPLE P31 Q42 Q5 normal
With -i, one Edge per line; a line may start with #label, and later lines
may refer to that label: it becomes the Edge Id the line gets.`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			list, _ := cmd.Flags().GetString("input")
			lines := []string{strings.Join(args[1:], " ")}
			if list != "" {
				var err error
				if lines, err = a.readLines(list); err != nil {
					return err
				}
			}
			return a.add(args[0], lines)
		},
	}
	cmd.Flags().StringP("input", "i", "", "file with one Edge per line (- for standard input)")
	return cmd
}

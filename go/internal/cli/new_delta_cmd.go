//ff:func feature=cli type=command control=sequence
//ff:what newDeltaCmd — geul delta: write the DELTA document that turns one version into another
package cli

import "github.com/spf13/cobra"

// newDeltaCmd is `geul delta <문서>@v1 @v2 -o <파일>`.
func newDeltaCmd(a *app) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delta DOCUMENT@V1 @V2 -o FILE",
		Short: "Write the diff document (DELTA-BASE, DELTA-REMOVE) that rebuilds V2 from V1",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			out, _ := cmd.Flags().GetString("output")
			return a.delta(args[0], args[1], out)
		},
	}
	cmd.Flags().StringP("output", "o", "", "file to write")
	_ = cmd.MarkFlagRequired("output")
	return cmd
}

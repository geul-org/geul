//ff:func feature=cli type=command control=sequence
//ff:what newNewCmd — geul new: start a document draft with DOC-ID and SPEC-VERSION
package cli

import "github.com/spf13/cobra"

// newNewCmd is `geul new [<Issuer>/]<문서 ID> -o <파일>` (git init).
func newNewCmd(a *app) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new [ISSUER/]DOCUMENT-ID -o FILE",
		Short: "Start a draft: SPEC-VERSION, DOC-ID (and ISSUER-KEY for a local self-issued Issuer), empty BODY",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			out, _ := cmd.Flags().GetString("output")
			force, _ := cmd.Flags().GetBool("force")
			return a.newDraft(args[0], out, force)
		},
	}
	cmd.Flags().StringP("output", "o", "", "draft file to write")
	cmd.Flags().Bool("force", false, "replace an existing file")
	_ = cmd.MarkFlagRequired("output")
	return cmd
}

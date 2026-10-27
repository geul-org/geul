//ff:func feature=cli type=command control=sequence
//ff:what newTombstoneCmd — geul tombstone: delete a version with a reason
package cli

import "github.com/spf13/cobra"

// newTombstoneCmd is `geul tombstone <문서>@<버전> --reason <n>`.
func newTombstoneCmd(a *app) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tombstone DOCUMENT@VERSION --reason N",
		Short: "Append a TOMBSTONE with its reason (LICENSE-GEUL-ORG 제5조) and remove the version's files",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			reason, _ := cmd.Flags().GetUint16("reason")
			k, err := a.versionKey(args[0])
			if err != nil {
				return err
			}
			d, _ := a.dir()
			return d.Tombstone(k.Issuer, k.Doc, k.Version, reason)
		},
	}
	cmd.Flags().Uint16("reason", 0, "reason category (사유 범주)")
	_ = cmd.MarkFlagRequired("reason")
	return cmd
}

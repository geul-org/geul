//ff:func feature=cli type=command control=sequence
//ff:what newIssuerCmd — geul issuer: self-issued 16-word Issuers and their local keys
package cli

import "github.com/spf13/cobra"

// newIssuerCmd groups `issuer new`, `issuer ls` and `issuer show`. Only
// self-issued (16-word) Issuers are made here; geul.org registration is a
// later task (README 차후 과제).
func newIssuerCmd(a *app) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issuer",
		Short: "Make and list self-issued Issuers (16-word IDs from an Ed25519 key kept outside the data folder)",
	}
	newCmd := &cobra.Command{
		Use:   "new [--name NAME]",
		Short: "Make an Ed25519 key and its 16-word Issuer ID; the key is saved under the user's config folder",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			return a.issuerNew(name)
		},
	}
	newCmd.Flags().String("name", "default", "local name of the Issuer (a-z, 0-9, -)")
	cmd.AddCommand(newCmd, &cobra.Command{
		Use:   "ls",
		Short: "List the local Issuers: name and ID",
		Args:  cobra.NoArgs,
		RunE:  func(cmd *cobra.Command, args []string) error { return a.issuerLs() },
	}, &cobra.Command{
		Use:   "show NAME|ID",
		Short: "Print a local Issuer's ID, algorithm and public key",
		Args:  cobra.ExactArgs(1),
		RunE:  func(cmd *cobra.Command, args []string) error { return a.issuerShow(args[0]) },
	})
	return cmd
}

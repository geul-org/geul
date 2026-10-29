//ff:func feature=cli type=command control=sequence
//ff:what newPackCmd — geul pack: convert a directory archive to pack sets and back, verify and compact them
package cli

import "github.com/spf13/cobra"

// newPackCmd groups `pack build`, `pack unpack`, `pack verify` and `pack compact`.
func newPackCmd(a *app) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pack",
		Short: "Pack sets: convert a directory archive (build) and back (unpack), verify, compact",
		Long: `A pack archive keeps .well-known/geul/ as immutable content-named objects
(m/<sha256>.gpm manifests, o/<sha256>.gpx indexes, o/<sha256>.gpk data packs) and
latest.txt, one set per Issuer (<Issuer>/latest.txt) or one at the top. Commands
that write to a data folder holding a pack archive add a generation per changed
set after the command succeeds.`,
	}
	var out string
	var single bool
	build := &cobra.Command{
		Use:   "build -o OUT [DIR]",
		Short: "Pack the directory archive DIR (the data folder by default) into OUT/.well-known/geul/",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			src, err := a.target(args)
			if err != nil {
				return err
			}
			return packBuild(src, out, single, a.out)
		},
	}
	build.Flags().StringVarP(&out, "out", "o", "", "output data folder")
	build.Flags().BoolVar(&single, "single", false, "one set at the top instead of one per Issuer")
	build.MarkFlagRequired("out")
	unpack := &cobra.Command{
		Use:   "unpack -o OUT [DIR]",
		Short: "Write the pack archive DIR (the data folder by default) as a directory tree under OUT",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			src, err := a.target(args)
			if err != nil {
				return err
			}
			return packUnpack(src, out, a.out)
		},
	}
	unpack.Flags().StringVarP(&out, "out", "o", "", "output data folder")
	unpack.MarkFlagRequired("out")
	verify := &cobra.Command{
		Use:   "verify [DIR]",
		Short: "Check every object's SHA-256 and that each set's index matches its inventories",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			src, err := a.target(args)
			if err != nil {
				return err
			}
			return packVerify(src, a.out)
		},
	}
	compact := &cobra.Command{
		Use:   "compact [DIR]",
		Short: "Rewrite each set's live records into fresh objects (older objects are left for removal)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			src, err := a.target(args)
			if err != nil {
				return err
			}
			return packCompact(src, a.out)
		},
	}
	cmd.AddCommand(build, unpack, verify, compact)
	return cmd
}

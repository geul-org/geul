//ff:func feature=cli type=helper control=sequence
//ff:what app.load — read the settings in effect, the persistent flags taking precedence
package cli

import (
	"github.com/geul-org/geul/go/internal/config"
	"github.com/spf13/cobra"
)

// load fills a.cfg from the --data, --format and --config flags of cmd and
// a.env (config.Load).
func (a *app) load(cmd *cobra.Command) error {
	flag := func(name string) string { v, _ := cmd.Flags().GetString(name); return v }
	a.env.Flags = map[string]string{"GEULDATA": flag("data"), "GEULFORMAT": flag("format"), "GEULENV": flag("config")}
	cfg, err := config.Load(a.env)
	a.cfg = cfg
	return err
}

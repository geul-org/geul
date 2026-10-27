//ff:func feature=cli type=helper control=sequence
//ff:what app.target — the file or folder a command was given, or the data folder when none
package cli

// target is args' one path, or the data folder (GEULDATA) when args is empty,
// for serve and build.
func (a *app) target(args []string) (string, error) {
	if len(args) == 1 {
		return args[0], nil
	}
	d, err := a.dir()
	return d.Root, err
}

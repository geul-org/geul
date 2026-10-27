//ff:func feature=config type=helper control=selection
//ff:what sourceOf — where the settings file path came from: flag, environment or default
package config

// sourceOf names the source of GEULENV.
func sourceOf(flag, env string) string {
	switch {
	case flag != "":
		return "flag"
	case env != "":
		return "env"
	default:
		return "default"
	}
}

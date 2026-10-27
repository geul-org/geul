//ff:type feature=config type=model
//ff:what Config — the settings in effect, each with its source, and the settings file they were read from
package config

// Config holds every key of Keys (Get, Source). File is the settings file
// path in effect (GEULENV).
type Config struct {
	values map[string]Value
	File   string
}

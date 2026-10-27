//ff:func feature=config type=helper control=iteration dimension=1
//ff:what Load — every setting by precedence: flag, environment variable, settings file, then the found data folder or default
package config

import (
	"fmt"
	"path/filepath"
)

// Load resolves each key of Keys from e: a flag, else the environment
// variable, else the settings file (GEULENV, else `<ConfigDir>/geul/env`),
// else defaults. GEULFORMAT must be text or json.
func Load(e Env) (Config, error) {
	c := Config{values: map[string]Value{}, File: firstOf(e.Flags["GEULENV"], e.Getenv("GEULENV"), filepath.Join(e.ConfigDir, "geul", "env"))}
	file, err := readFile(c.File)
	if err != nil {
		return Config{}, err
	}
	def := defaults(e.Cwd, e.CacheDir)
	for _, k := range Keys {
		c.values[k] = pick(e.Flags[k], e.Getenv(k), file[k], def[k])
	}
	c.values["GEULENV"] = Value{Value: c.File, Source: sourceOf(e.Flags["GEULENV"], e.Getenv("GEULENV"))}
	if f := c.Get("GEULFORMAT"); f != "text" && f != "json" {
		return Config{}, fmt.Errorf("GEULFORMAT %q is neither text nor json", f)
	}
	return c, nil
}

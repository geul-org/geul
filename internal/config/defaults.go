//ff:func feature=config type=helper control=sequence
//ff:what defaults — the value of each key when no flag, variable or file sets it
package config

import "path/filepath"

// defaults gives GEULDATA the folder found above cwd (Discover, source
// "found"), GEULCACHE `<cache>/geul`, GEULREMOTE geul.org and GEULFORMAT
// text. Other keys have no default here (resolve applies its own).
func defaults(cwd, cache string) map[string]Value {
	out := map[string]Value{
		"GEULREMOTE": {Value: "https://geul.org", Source: "default"},
		"GEULFORMAT": {Value: "text", Source: "default"},
	}
	if cache != "" {
		out["GEULCACHE"] = Value{Value: filepath.Join(cache, "geul"), Source: "default"}
	}
	if d := Discover(cwd); d != "" {
		out["GEULDATA"] = Value{Value: d, Source: "found"}
	}
	return out
}

//ff:func feature=config type=helper control=selection
//ff:what pick — a setting from the first source that has it: flag, environment, file, default
package config

// pick returns the first non-empty of flag, env and file with its source,
// else def.
func pick(flag, env, file string, def Value) Value {
	switch {
	case flag != "":
		return Value{Value: flag, Source: "flag"}
	case env != "":
		return Value{Value: env, Source: "env"}
	case file != "":
		return Value{Value: file, Source: "file"}
	default:
		return def
	}
}

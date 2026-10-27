//ff:type feature=config type=model
//ff:what Value — one setting's value and where it came from
package config

// Value is a setting and its Source: "flag", "env", "file", "found" (the data
// folder found above the working directory), "default", or "" when unset.
type Value struct {
	Value  string `json:"value"`
	Source string `json:"source,omitempty"`
}

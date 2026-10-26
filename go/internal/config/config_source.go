//ff:func feature=config type=helper control=sequence
//ff:what Config.Source — where a setting's value came from
package config

// Source returns the source of key's value (Value).
func (c Config) Source(key string) string {
	return c.values[key].Source
}

//ff:func feature=config type=helper control=sequence
//ff:what Config.Get — a setting's value ("" when unset)
package config

// Get returns the value of key.
func (c Config) Get(key string) string {
	return c.values[key].Value
}

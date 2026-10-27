//ff:func feature=config type=helper control=iteration dimension=1
//ff:what firstOf — the first non-empty string
package config

// firstOf returns the first of ss that is not empty, or "".
func firstOf(ss ...string) string {
	for _, s := range ss {
		if s != "" {
			return s
		}
	}
	return ""
}

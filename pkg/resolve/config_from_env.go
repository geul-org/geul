//ff:func feature=resolve type=decoder control=sequence
//ff:what ConfigFromEnv — Config from GEULPROXY (default geul.org) and GEULPRIVATE
package resolve

// ConfigFromEnv builds a Config from the environment read through getenv.
// An empty GEULPROXY means DefaultProxy, so geul.org is used unless a custom
// source list (another server, a `file://` tree) is set.
func ConfigFromEnv(getenv func(string) string) (Config, error) {
	proxy := getenv("GEULPROXY")
	if proxy == "" {
		proxy = DefaultProxy
	}
	entries, err := parseProxy(proxy)
	if err != nil {
		return Config{}, err
	}
	private, err := parsePatterns(getenv("GEULPRIVATE"))
	if err != nil {
		return Config{}, err
	}
	return Config{Proxy: entries, Private: private}, nil
}

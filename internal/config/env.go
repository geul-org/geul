//ff:type feature=config type=model
//ff:what Env — what Load reads settings from: flags, the environment, the working directory and the user folders
package config

// Env is Load's input. Flags are the values given on the command line by key
// (empty means not given). Getenv reads the environment. Cwd is where the data
// folder search starts. ConfigDir and CacheDir are the user's folders
// (os.UserConfigDir, os.UserCacheDir); the default settings file is
// `<ConfigDir>/geul/env`.
type Env struct {
	Flags     map[string]string
	Getenv    func(string) string
	Cwd       string
	ConfigDir string
	CacheDir  string
}

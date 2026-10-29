//ff:type feature=config type=model
//ff:what Keys — the settings geul reads: GEUL-ORG 위치 해석 names and the CLI's own, in display order
package config

// Keys are every setting, environment variable and settings file key alike
// (GEUL-ORG 위치 해석: same names, the environment wins). GEULDATA is the data
// folder (the root holding `.well-known/geul/`), GEULENV the settings file,
// GEULCACHE the cache of fetched documents, GEULISSUER the default Issuer of
// `geul new`, GEULREMOTE where `push` sends, GEULFORMAT the output (text or
// json), GEULSTORE how a new data folder stores its archive (pack sets by
// default; dir for the directory tree), GEULPACK pack sets to read before
// GEULPROXY (comma-separated https:// bases or local set directories). The
// rest are resolve's (GEULPROXY …).
var Keys = []string{"GEULDATA", "GEULENV", "GEULCACHE", "GEULISSUER", "GEULREMOTE", "GEULFORMAT", "GEULSTORE", "GEULPACK", "GEULPROXY", "GEULPRIVATE", "GEULNOLOG", "GEULLOG", "GEULINSECURE", "GEULWORK"}

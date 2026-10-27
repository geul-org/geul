//ff:type feature=resolve type=resolver
//ff:what Resolver — fetches, verifies and links documents according to a Config, optionally as of a point in time
package resolve

import "time"

// Resolver resolves documents with Config. Fetching is always an explicit call;
// decoding a document never fetches. AsOf, when set, binds latest (version 0)
// imports to the version in effect at that instant instead of the newest one
// (OVERVIEW 계보 개정).
type Resolver struct {
	Config Config
	AsOf   time.Time
}

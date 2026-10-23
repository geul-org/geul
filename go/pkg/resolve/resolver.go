//ff:type feature=resolve type=resolver
//ff:what Resolver — fetches, verifies and links documents according to a Config
package resolve

// Resolver resolves documents with Config. Fetching is always an explicit call;
// decoding a document never fetches.
type Resolver struct {
	Config Config
}

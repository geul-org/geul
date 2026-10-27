//ff:type feature=resolve type=resolver
//ff:what Registry — the geul.org registry source; its lookup API is not fixed yet
package resolve

// Registry is the `registry` GEULPROXY entry: look up an Issuer's base address
// in the geul.org registry and fetch from there. The registry API is not fixed
// yet (GEUL-ORG 미결정), so every lookup reports ErrNotFound for now.
type Registry struct{}

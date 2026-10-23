//ff:type feature=resolve type=resolver
//ff:what Dir — a Source and publish target over a local `<Root>/.well-known/geul/` tree
package resolve

// Dir is a local directory laid out like the public address space:
// `<Root>/.well-known/geul/<path>`. Dumps, mirrors and caches share the layout.
// A version list path (ending in `/`) is the `versions` file in that directory.
type Dir struct {
	Root string
}

//ff:type feature=resolve type=resolver
//ff:what Dir — a Source and publish target over a local archive: `<Root>/.well-known/geul/` as a directory tree, or any Store
package resolve

// Dir is an archive laid out like the public address space. Store nil means
// the directory tree `<Root>/.well-known/geul/<path>` (dumps, mirrors and
// caches share it); a Store (pack sets) replaces the tree, and Root is then
// only the data folder's name.
type Dir struct {
	Root  string
	Store Store
}

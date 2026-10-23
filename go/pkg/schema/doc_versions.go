//ff:type feature=schema type=model
//ff:what docVersions — one document's state while reading a version list
package schema

// docVersions is what a version list has said so far about one document: the
// last published version ID and the content hashes already published.
type docVersions struct {
	last   uint32
	hashes map[[32]byte]bool
}

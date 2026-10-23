//ff:type feature=resolve type=model
//ff:what VersionEntry — one version list line: version ID, content hash, deprecated flag
package resolve

// VersionEntry is one entry of a document's version list (GEUL-ORG 공표 주소):
// the version ID, the content hash that version must have, and whether it is
// deprecated. A deprecated version is still served but never picked as latest.
type VersionEntry struct {
	Version    uint32
	Hash       [32]byte
	Deprecated bool
}

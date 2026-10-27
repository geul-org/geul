//ff:type feature=resolve type=model
//ff:what VersionEntry — one version of a document: version ID, content hash, deprecated and deleted flags
package resolve

// VersionEntry is one version as its version list states it (edge/PUBLICATION):
// the version ID, the content hash that version must have, and whether a later
// record deprecated or deleted (TOMBSTONE) it. A deprecated version is still
// served but never picked as latest; a deleted one is not served.
type VersionEntry struct {
	Version    uint32
	Hash       [32]byte
	Deprecated bool
	Deleted    bool
}

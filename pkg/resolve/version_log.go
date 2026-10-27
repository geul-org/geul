//ff:type feature=resolve type=model
//ff:what versionLog — a decoded version list document: its versions, its BODY records and its content hash
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// versionLog is one version list document: the versions it states, its BODY
// records as written (a new list keeps them and appends), and its content hash
// (the next list's PARENT) and its own PARENT, the previous list's hash, when
// HasParent. Present is false when the document has no list yet.
type versionLog struct {
	Entries   []VersionEntry
	Records   []core.Record
	Hash      [32]byte
	Parent    [32]byte
	HasParent bool
	Next      uint64
	Present   bool
}

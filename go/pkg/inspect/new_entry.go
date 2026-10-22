//ff:func feature=inspect type=decoder control=sequence
//ff:what newEntry — a bare Entry: the record, its encoding and its core spec name
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// newEntry returns the uninterpreted Entry of r: the record, its words (a
// decoded record always re-encodes) and the core name of its codec.
func newEntry(r core.Record) Entry {
	ws, _ := core.EncodeRecord(r)
	return Entry{Record: r, Words: ws, Name: core.CodecName(r.Kind, r.Codec)}
}

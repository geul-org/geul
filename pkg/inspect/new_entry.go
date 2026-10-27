//ff:func feature=inspect type=decoder control=sequence
//ff:what newEntry — a bare Entry: the record, its encoding and its core spec name (a slot marker by its marker name)
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// newEntry returns the uninterpreted Entry of r: the record, its words (a
// decoded record always re-encodes) and its core name (core.RecordName: a
// Value ISSUER format by its format name, else the codec name), or the name of the
// slot marker it is (core.SlotMarkerName; misplaced as a record, which schema
// reports).
func newEntry(r core.Record) Entry {
	ws, _ := core.EncodeRecord(r)
	name := core.RecordName(r)
	if core.IsSlotMarker(r) {
		name = core.SlotMarkerName(r.Codec)
	}
	return Entry{Record: r, Words: ws, Name: name}
}

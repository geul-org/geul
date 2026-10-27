//ff:func feature=core type=helper control=sequence
//ff:what IsSlotMarker — whether a decoded record is a slot marker: Reference Codec 0000, one word (no payload, no ISSUER extension)
package core

// IsSlotMarker reports whether r is a one-word record of the reserved
// Reference Codec 0000 (slot marker family, TAG WORD 0x0000~0x0007): its low
// 3 bits are the marker number (SlotMarkerName). A Codec 0000 Reference with a
// payload is not a marker; it is reserved.
func IsSlotMarker(r Record) bool {
	return r.Kind == KindReference && r.Codec>>3 == 0 && len(r.Payload) == 0 && len(r.Issuer) == 0
}

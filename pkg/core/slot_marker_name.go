//ff:func feature=core type=helper control=selection
//ff:what SlotMarkerName — the name of a slot marker number (low 3 bits of Reference Codec 0000): ABSENT, UNKNOWN-VALUE, NO-VALUE, else ""
package core

// SlotMarkerName names marker number m (GEUL reference/slot-marker): 0 ABSENT
// (적지 않음), 1 UNKNOWN-VALUE (알 수 없음), 2 NO-VALUE (없음). Numbers
// 3~7 are reserved and return "". Only the low 3 bits are read. The number
// is not a Quantifier.
func SlotMarkerName(m uint8) string {
	switch m & 7 {
	case 0:
		return "ABSENT"
	case 1:
		return "UNKNOWN-VALUE"
	case 2:
		return "NO-VALUE"
	default:
		return ""
	}
}

//ff:type feature=core type=model
//ff:what Absent, UnknownValue, NoValue — the slot markers: one-word Reference records of the reserved Codec 0000 (TAG WORD 0x0000~0x0002)
package core

// The slot markers (GEUL reference/slot-marker) are one-word records of the
// reserved Reference Codec 0000: TAG WORD 0000 0000 0000 0MMM, where the low
// 3 bits MMM are a marker number, not a Quantifier. A marker points at
// nothing; it fills a record field whose CODEC list names that exact WORD
// (CodecMatches), never a HEAD or BODY record position, an Edge Id, a field
// without CODECs or a wildcard. 0x0003~0x0007 are reserved.
const (
	// Absent (적지 않음) fills an optional field left out. It says nothing
	// about the world.
	Absent Word = 0x0000
	// UnknownValue (알 수 없음) says a value exists but is not known
	// (Wikidata somevalue): ∃x, x unidentified.
	UnknownValue Word = 0x0001
	// NoValue (없음) says no value exists (Wikidata novalue): ¬∃x, closed
	// within the statement's qualifiers.
	NoValue Word = 0x0002
)

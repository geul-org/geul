//ff:type feature=core type=model
//ff:what Record — one GEUL record: Kind, Codec, optional ISSUER extension and Payload
package core

// Record is one decoded GEUL record (GEUL t002 README). Codec is the 7-bit
// codec of the TAG WORD. When Codec is an ISSUER code (0x7C-0x7F on a Value or
// Edge), Issuer holds the 1/2/4/16-word Issuer and Format the 1-word format
// number of the Codec extension TAG; otherwise both are empty. Payload is
// everything after the TAG WORD and its extensions. The Length field is not
// stored: it is derived from the encoded size.
type Record struct {
	Kind    Kind
	Codec   uint8
	Issuer  []Word
	Format  uint16
	Payload []Word
}

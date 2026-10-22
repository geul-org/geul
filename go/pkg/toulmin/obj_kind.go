//ff:type feature=toulmin type=model
//ff:what ObjKind — the 2-bit discriminator of a Claim object slot (REF | integer | SPAN)
package toulmin

// ObjKind classifies the object slot of a Claim block. The Claim carries no kind
// word; the 2-bit kind rides in the NODE header word (bits 8-7) so the decoder
// knows the object's form before reading the Claim (GEUL t002 TOULMIN NODE).
type ObjKind uint16

const (
	// ObjTID is a REF object (the default).
	ObjTID ObjKind = iota
	// ObjNum is an UNSIGNED-INTEGER Value object.
	ObjNum
	// ObjSpan is a SPAN object — the descriptive-claim fallback anchor.
	ObjSpan
)

//ff:type feature=core type=model
//ff:what FieldDecl — a FIELD Edge: Edge Id, DECLARE it belongs to, width (0 = record, n = raw words), optional expected TAG
package core

// FieldDecl is a FIELD Edge (GEUL t002 docs/edge/FIELD). Declare is the Edge
// Id of the DECLARE it belongs to. Width 0 is a record field read by its own
// TAG Length; n is a raw field of exactly n words; FieldRest is the raw rest
// of the record. Expect, when HasExpect, is the Kind·Codec a record field must
// have.
type FieldDecl struct {
	EdgeID    uint64
	Declare   uint64
	Width     Word
	Expect    Word
	HasExpect bool
}

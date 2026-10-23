//ff:type feature=core type=model
//ff:what Declare — a DECLARE Edge: Edge Id, target Kind·Codec, Issuer and format of an ISSUER target, FIELD Edge Ids
package core

// Declare is a DECLARE Edge (GEUL t002 docs/edge/DECLARE). Target is the
// declared codec's TAG WORD with only the Kind and Codec bits set. Issuer and
// Format are set for an ISSUER target. Fields lists the FIELD Edge Ids in order.
type Declare struct {
	EdgeID uint64
	Target Word
	Issuer []Word
	Format Word
	Fields []uint64
}

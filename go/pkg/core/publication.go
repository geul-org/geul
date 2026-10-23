//ff:type feature=core type=model
//ff:what Publication — one version list record: PUBLICATION, DEPRECATION or TOMBSTONE
package core

// Publication is one version list record (GEUL t002 docs/edge/PUBLICATION).
// Codec says which: EdgePublication carries Hash, EdgeTombstone carries Reason,
// EdgeDeprecation carries neither. Doc is the 4-word document ID; the Issuer is
// the version list document's own.
type Publication struct {
	Codec   uint8
	EdgeID  uint64
	Doc     []Word
	Version uint32
	Hash    [32]byte
	Reason  uint16
}

package core

// Edge codecs the core reads without a codebook (GEUL t002 OVERVIEW 서술 Edge,
// docs/edge/PUBLICATION, edge/BIND, edge/DECLARE): the version list records and
// the codec spec vocabulary (GEUL-ORG 0층).
const (
	// EdgeBind binds a machine-matched name to a Subject Reference (서술 0x03).
	// Version list documents name documents with it (edge/BIND, edge/PUBLICATION).
	EdgeBind uint8 = 0x03
	// EdgePublication records one published version: Edge Id, document ID, version ID, content hash.
	EdgePublication uint8 = 0x18
	// EdgeDeprecation deprecates a published version: Edge Id, document ID, version ID.
	EdgeDeprecation uint8 = 0x19
	// EdgeTombstone records a deleted version: Edge Id, document ID, version ID, reason.
	EdgeTombstone uint8 = 0x1A
	// EdgeDeclare declares one codec's payload: Edge Id, target, FIELD list (edge/DECLARE).
	EdgeDeclare uint8 = 0x1B
	// EdgeField declares one payload field: Edge Id, DECLARE, width (edge/FIELD).
	EdgeField uint8 = 0x1C
	// EdgeEnum names one code of a raw field: Edge Id, field, code (edge/ENUM).
	EdgeEnum uint8 = 0x1D
	// EdgeCodec names one Kind·Codec a field takes: Edge Id, field, target word (edge/CODEC).
	EdgeCodec uint8 = 0x1E
)

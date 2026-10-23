package core

// Edge codecs the core reads without a codebook (GEUL t002 OVERVIEW 서술 Edge,
// docs/edge/PUBLICATION, edge/BIND): the version list records (GEUL-ORG 0층).
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
)

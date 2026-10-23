package core

// Edge codecs the core reads without a codebook (GEUL t002 OVERVIEW 서술 Edge,
// docs/edge/PUBLICATION): the version list records (GEUL-ORG 0층).
const (
	// EdgePublication records one published version: Edge Id, document ID, version ID, content hash.
	EdgePublication uint8 = 0x18
	// EdgeDeprecation deprecates a published version: Edge Id, document ID, version ID.
	EdgeDeprecation uint8 = 0x19
	// EdgeTombstone records a deleted version: Edge Id, document ID, version ID, reason.
	EdgeTombstone uint8 = 0x1A
)

//ff:type feature=pack type=model
//ff:what Stored — the uncompressed Codec: zstd frames of raw blocks, checksummed; records stay raw
package pack

// Stored writes zstd frames whose blocks are raw (Block_Type 0) with a
// Content_Checksum, so blocks and inventories keep their integrity check.
// Records never shrink under Stored, so writers keep them raw (encoding 0).
type Stored struct{}

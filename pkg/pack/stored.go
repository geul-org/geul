//ff:type feature=pack type=model
//ff:what Stored — the built-in Codec: zstd frames of raw (uncompressed) blocks, checksummed, no outside module
package pack

// Stored writes zstd frames whose blocks are raw (Block_Type 0) with a
// Content_Checksum, so blocks and inventories keep their integrity check, and
// reads raw and RLE blocks. A compressed block is ErrUnsupported. Records
// never shrink under Stored, so writers keep them raw (encoding 0).
type Stored struct{}

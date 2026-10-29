//ff:type feature=pack type=model
//ff:what Options — how to open a generation: pinned manifest or latest.txt, rollback floor, codec, block cache size
package pack

// Options for Open. Manifest pins a generation by its SHA-256 hex (latest.txt
// is then not read). MinGeneration rejects an older latest.txt (rollback).
// Codec nil is Stored. CacheBlocks 0 is DefaultCacheBlocks.
type Options struct {
	Manifest      string
	MinGeneration uint64
	Codec         Codec
	CacheBlocks   int
}

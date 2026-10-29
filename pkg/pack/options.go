//ff:type feature=pack type=model
//ff:what Options — how to open a generation: pinned manifest or latest.txt, rollback floor, block cache size
package pack

// Options for Open. Manifest pins a generation by its SHA-256 hex (latest.txt
// is then not read). MinGeneration rejects an older latest.txt (rollback).
// CacheBlocks 0 is DefaultCacheBlocks. No codec is needed to read.
type Options struct {
	Manifest      string
	MinGeneration uint64
	CacheBlocks   int
}

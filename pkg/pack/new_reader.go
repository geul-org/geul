//ff:func feature=pack type=resolver control=sequence
//ff:what newReader — a Reader over a parsed manifest with defaults filled in
package pack

import "github.com/klauspost/compress/zstd"

// newReader builds the Reader state; CacheBlocks 0 is the default.
func newReader(s Store, o Options, m *manifest, name string) *Reader {
	r := &Reader{s: s, m: m, name: name, inum: issuerIndex(m.issuers), blocks: map[Loc]block{},
		dicts: map[uint32][]byte{}, decs: map[uint32]*zstd.Decoder{}, limit: o.CacheBlocks}
	if r.limit <= 0 {
		r.limit = DefaultCacheBlocks
	}
	return r
}

//ff:func feature=pack type=resolver control=sequence
//ff:what newReader — a Reader over a parsed manifest with defaults filled in
package pack

// newReader builds the Reader state; Codec nil is Stored, CacheBlocks 0 the default.
func newReader(s Store, o Options, m *manifest, name string) *Reader {
	r := &Reader{s: s, codec: o.Codec, m: m, name: name, inum: issuerIndex(m.issuers),
		blocks: map[Loc]block{}, dicts: map[uint32][]byte{}, limit: o.CacheBlocks}
	if r.codec == nil {
		r.codec = Stored{}
	}
	if r.limit <= 0 {
		r.limit = DefaultCacheBlocks
	}
	return r
}

//ff:func feature=pack type=writer control=sequence
//ff:what Stream.hFlush — write the H leaf when it is full, returning the heads with its first key and an empty leaf
package pack

// hFlush writes leaf as an H leaf block once it holds HLeaf entries and adds
// its head; otherwise it returns both unchanged.
func (s *Stream) hFlush(heads, leaf []streamKV) ([]streamKV, []streamKV, error) {
	if len(leaf) < s.o.Build.HLeaf {
		return heads, leaf, nil
	}
	loc, err := s.addBlock(blockOf(blkHLeaf, leaf))
	if err != nil {
		return nil, nil, err
	}
	return append(heads, streamKV{Key: leaf[0].Key, Loc: loc}), leaf[:0], nil
}

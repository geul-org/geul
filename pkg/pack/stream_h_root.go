//ff:func feature=pack type=writer control=sequence
//ff:what Stream.hRoot — the H layer from the written leaf heads and the last leaf: that leaf alone is the root, else it is written and the tree built
package pack

// hRoot finishes the H layer of n keys: with no written leaf the pending one
// is the root; otherwise it is written too and tree groups the heads.
func (s *Stream) hRoot(heads, leaf []streamKV, n uint64) ([]layer, error) {
	if n == 0 {
		return nil, nil
	}
	if len(heads) == 0 {
		return s.layerOf(blockOf(blkHLeaf, leaf), n)
	}
	loc, err := s.addBlock(blockOf(blkHLeaf, leaf))
	if err != nil {
		return nil, err
	}
	root, err := s.tree(append(heads, streamKV{Key: leaf[0].Key, Loc: loc}), blkHInner, s.o.Build.HFanout, s.o.Build.HRoot)
	if err != nil {
		return nil, err
	}
	return s.layerOf(root, n)
}

//ff:func feature=pack type=writer control=sequence
//ff:what Stream.pLayer — the P layer: the pending leaf alone as root, or written and one root over every leaf (depth 2)
package pack

// pLayer finishes the P layer: with no written leaf the pending one is the
// root (no records: no layer); otherwise the pending leaf is written and one
// internal root holds every leaf's head, as Write does for P (§9.2, depth 2).
func (s *Stream) pLayer() ([]layer, error) {
	n := uint64(s.st.Count)
	if n == 0 {
		return nil, nil
	}
	if len(s.st.PHeads) == 0 {
		return s.layerOf(blockOf(blkPLeaf, s.st.PLeaf), n)
	}
	if err := s.flushLeaf(); err != nil {
		return nil, err
	}
	return s.layerOf(blockOf(blkPInner, s.st.PHeads), n)
}

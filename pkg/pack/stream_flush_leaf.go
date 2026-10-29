//ff:func feature=pack type=writer control=sequence
//ff:what Stream.flushLeaf — write the filled P leaf block and keep only its first key and location
package pack

// flushLeaf writes the pending P leaf (§4.7 kind 1) and records its head for
// the root. It runs when the next record arrives after a full leaf, so a set
// with a single leaf keeps it unwritten as its root, as Write does.
func (s *Stream) flushLeaf() error {
	loc, err := s.addBlock(blockOf(blkPLeaf, s.st.PLeaf))
	if err != nil {
		return err
	}
	s.st.PHeads = append(s.st.PHeads, streamKV{Key: s.st.PLeaf[0].Key, Loc: loc})
	s.st.PLeaf = s.st.PLeaf[:0]
	return nil
}

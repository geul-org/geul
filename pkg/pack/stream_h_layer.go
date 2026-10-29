//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Stream.hLayer — merge the H runs into leaf blocks (a repeated key keeps its last record) and build the H tree over them
package pack

// hLayer spills the held keys, merges every run in key order and writes H
// leaves of HLeaf entries (§4.7 kind 3). Of entries with the same key the
// last (hEntryLess) wins, as it would when the inventories are applied in
// order. A leaf is written when the next key arrives after it filled, so a
// single leaf stays unwritten as the root. The levels above are built by
// tree with HFanout and HRoot. No key gives no layer.
func (s *Stream) hLayer() ([]layer, error) {
	if err := s.spill(); err != nil {
		return nil, err
	}
	rs, err := openRuns(s.out, s.st.Runs)
	if err != nil {
		return nil, err
	}
	defer closeRuns(rs)
	var leaf, heads []streamKV
	var n uint64
	for {
		e, ok, err := nextMerged(rs)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		if k := string(e.key[:]); len(leaf) > 0 && leaf[len(leaf)-1].Key == k {
			leaf[len(leaf)-1].Loc = e.loc
			continue
		}
		if heads, leaf, err = s.hFlush(heads, leaf); err != nil {
			return nil, err
		}
		leaf, n = append(leaf, streamKV{Key: string(e.key[:]), Loc: e.loc}), n+1
	}
	return s.hRoot(heads, leaf, n)
}

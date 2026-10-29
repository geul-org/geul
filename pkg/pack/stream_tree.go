//ff:func feature=pack type=writer control=iteration dimension=2
//ff:what Stream.tree — the root block over written children: grouped by fanout into written internal blocks while more than rootMax remain (rootMax 0: one root over all)
package pack

// tree builds the internal levels over heads (first key and location of each
// written child block), as Write's writeLayer does: while more than rootMax
// remain (never with rootMax 0) they are grouped fanout at a time into
// internal blocks of kind, which are written unless a level has only one,
// which is then the root. The root (not written) is returned.
func (s *Stream) tree(heads []streamKV, kind byte, fanout, rootMax int) (block, error) {
	for rootMax != 0 && len(heads) > rootMax {
		if len(heads) <= fanout {
			return blockOf(kind, heads), nil
		}
		var next []streamKV
		for i := 0; i < len(heads); i += fanout {
			g := heads[i:min(i+fanout, len(heads))]
			loc, err := s.addBlock(blockOf(kind, g))
			if err != nil {
				return block{}, err
			}
			next = append(next, streamKV{Key: g[0].Key, Loc: loc})
		}
		heads = next
	}
	return blockOf(kind, heads), nil
}

//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what parents — internal blocks over written child blocks: each child's first key and location, fanout per block
package pack

// parents groups children (written at locs) into internal blocks of kind.
func parents(bs []block, locs []Loc, kind byte, fanout int) []block {
	var out []block
	for i := range bs {
		if i%fanout == 0 {
			out = append(out, block{kind: kind})
		}
		p := &out[len(out)-1]
		p.keys, p.locs = append(p.keys, bs[i].keys[0]), append(p.locs, locs[i])
	}
	return out
}

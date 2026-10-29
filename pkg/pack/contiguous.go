//ff:func feature=pack type=helper control=sequence
//ff:what contiguous — whether l starts where p ends in the same object (C5 = 0 is allowed)
package pack

// contiguous reports whether l directly follows p in the same object and p is
// not a deletion marker.
func contiguous(p, l Loc) bool {
	return p.Enc != encGone && p.Obj == l.Obj && p.Off+p.Len == l.Off
}

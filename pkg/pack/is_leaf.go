//ff:func feature=pack type=helper control=sequence
//ff:what isLeaf — whether a block kind is a P or H leaf
package pack

// isLeaf reports whether kind is a leaf block kind.
func isLeaf(kind byte) bool {
	return kind == blkPLeaf || kind == blkHLeaf
}

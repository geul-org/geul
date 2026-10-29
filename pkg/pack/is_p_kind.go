//ff:func feature=pack type=helper control=sequence
//ff:what isPKind — whether a block kind belongs to the P space
package pack

// isPKind reports whether kind is a P leaf or P internal block.
func isPKind(kind byte) bool {
	return kind == blkPLeaf || kind == blkPInner
}

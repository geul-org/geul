//ff:func feature=pack type=decoder control=sequence
//ff:what decodeBlock — parse a directory block, rejecting unknown kinds, bad columns and trailing bytes
package pack

// decodeBlock parses the byte form of a directory block (§4.7).
func decodeBlock(raw []byte) (block, error) {
	c := &cursor{b: raw}
	kind := c.u8()
	n := c.uvarint()
	if c.bad || kind < blkPLeaf || kind > blkHInner || n > uint64(len(raw)) {
		return block{}, corrupt("directory block header")
	}
	b := block{kind: kind}
	var err error
	if isPKind(kind) {
		b.keys, err = readPKeys(c, int(n))
	} else {
		b.keys = readHKeys(c, int(n))
	}
	if err == nil {
		b.locs, err = readLocs(c, int(n), isLeaf(kind))
	}
	if err == nil {
		err = c.end("directory block")
	}
	if err == nil {
		err = checkAscending(b.keys)
	}
	return b, err
}

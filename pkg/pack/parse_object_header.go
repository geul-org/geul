//ff:func feature=pack type=decoder control=sequence
//ff:what parseObjectHeader — the 64-byte object head; magic, version, kind, reserved fields and inventory bounds checked
package pack

// parseObjectHeader decodes §4.4. Index objects have no inventory.
func parseObjectHeader(b []byte) (objectHeader, error) {
	c := &cursor{b: b}
	magic, ver, kind, res := string(c.take(8)), c.u16(), c.u16(), c.take(4)
	h := objectHeader{kind: byte(kind), size: c.u64(), invOff: c.u64(), invLen: c.u64(), count: c.u64()}
	res2 := c.take(16)
	if err := c.end("object header"); err != nil {
		return objectHeader{}, err
	}
	if magic != objectMagic || ver != formatVersion || kind != kindData && kind != kindIndex || !zeros(res) || !zeros(res2) {
		return objectHeader{}, corrupt("object header fields")
	}
	if kind == kindIndex && (h.invOff != 0 || h.invLen != 0) || kind == kindData && (h.invOff < headerSize || h.invOff > h.size || h.invLen > h.size-h.invOff) {
		return objectHeader{}, corrupt("object inventory bounds")
	}
	return h, nil
}

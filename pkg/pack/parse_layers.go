//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what parseLayers — P_LAYERS or H_LAYERS: u32 layers, each count, root encoding, root bytes (newest first)
package pack

// parseLayers decodes a layer section; roots are decoded later by the Reader.
func parseLayers(b []byte) ([]layer, error) {
	c := &cursor{b: b}
	n := c.u32()
	if c.bad || uint64(n)*16 > uint64(len(c.b)) {
		return nil, corrupt("layer section length")
	}
	out := make([]layer, n)
	for i := range out {
		cnt, enc, res := c.u64(), c.u8(), c.take(3)
		root := c.take(uint64(c.u32()))
		if c.bad || enc > encZstd || !zeros(res) {
			return nil, corrupt("layer entry")
		}
		out[i] = layer{count: cnt, enc: enc, root: root}
	}
	return out, c.end("layer section")
}

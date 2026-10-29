//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.writeLayer — one directory tree: leaves, internal levels until the root fits, root kept for the manifest
package pack

// writeLayer builds a layer over sorted kvs. While more than one block
// remains they are written and grouped: by fanout while more than rootMax
// remain, otherwise all under one root (rootMax 0: the root takes every leaf,
// depth 2). A single leaf is the root itself. Empty input makes no layer.
func (w *writer) writeLayer(kvs []kv, leaf byte, leafSize, fanout, rootMax int) ([]layer, error) {
	if len(kvs) == 0 {
		return nil, nil
	}
	bs := chunk(kvs, leaf, leafSize)
	for len(bs) > 1 {
		locs, err := w.writeBlocks(bs)
		if err != nil {
			return nil, err
		}
		n := fanout
		if rootMax == 0 || len(bs) <= rootMax {
			n = len(bs)
		}
		bs = parents(bs, locs, leaf+1, n)
	}
	root, err := w.o.Codec.Compress(encodeBlock(bs[0]), nil)
	if err != nil {
		return nil, err
	}
	return []layer{{count: uint64(len(kvs)), enc: encZstd, root: root}}, nil
}

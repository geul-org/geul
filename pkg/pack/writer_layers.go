//ff:func feature=pack type=writer control=sequence
//ff:what writer.layers — write the P layer (always one, rewritten) and the H layers (new delta first, or one merged)
package pack

// layers writes the index objects of the new generation and returns the P
// and H layer lists plus the NAMED table (§6 step 5).
func (w *writer) layers() ([]layer, []layer, []namedEntry, error) {
	ps, hs, ns := w.partition(w.newState())
	pl, err := w.writeLayer(ps, blkPLeaf, w.o.PLeaf, 0, 0)
	if err != nil {
		return nil, nil, nil, err
	}
	if !w.all {
		hs = w.hDelta()
	}
	hl, err := w.writeLayer(hs, blkHLeaf, w.o.HLeaf, w.o.HFanout, w.o.HRoot)
	if err == nil && w.cur != nil {
		err = w.closeObject(nil)
	}
	if err != nil {
		return nil, nil, nil, err
	}
	if !w.all {
		hl = append(hl, w.prev.m.hLayers...)
	}
	return pl, hl, ns, nil
}

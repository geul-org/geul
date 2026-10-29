//ff:func feature=pack type=writer control=sequence
//ff:what writer.open — load the previous generation: live state, Issuer table, OBJECTS and DICTS to carry and whether to rewrite all layers
package pack

// open prepares w from out's latest generation. Every layer is rewritten
// (and index objects dropped from OBJECTS) on the first generation, on Fresh,
// and when the H layers reached HLayers; otherwise OBJECTS only grows, so the
// kept H layers' object numbers stay valid.
func (w *writer) open() error {
	prev, err := openPrev(w.ctx, w.out)
	if err != nil {
		return err
	}
	w.prev, w.base, w.all = prev, map[string]Loc{}, true
	if prev == nil {
		w.inum = map[string]uint32{}
		return nil
	}
	w.res = BuildResult{Generation: prev.m.gen, Manifest: prev.name}
	w.issuers = append([]string(nil), prev.m.issuers...)
	w.inum = issuerIndex(w.issuers)
	if w.o.Fresh {
		return nil
	}
	if w.base, err = prev.entries(w.ctx); err != nil {
		return err
	}
	w.all = len(prev.m.hLayers) >= w.o.HLayers
	if w.all {
		w.objs, w.remap = keepData(prev.m.objects)
	} else {
		w.objs = append([]object(nil), prev.m.objects...)
	}
	w.dicts = w.carryDicts()
	return nil
}

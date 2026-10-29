//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.followCodec — the codec when none is given: Zstd if the set already has a dictionary or a compressed record, Stored otherwise
package pack

// followCodec keeps a set's codec across generations. The level is not
// recorded in the format, so a followed Zstd uses DefaultZstdLevel.
func (w *writer) followCodec() (Codec, error) {
	if w.prev == nil {
		return Stored{}, nil
	}
	if len(w.prev.m.dicts) > 0 {
		return &Zstd{}, nil
	}
	ls := w.base
	if w.o.Fresh {
		var err error
		if ls, err = w.prev.entries(w.ctx); err != nil {
			return nil, err
		}
	}
	for _, l := range ls {
		if l.Enc == encZstd {
			return &Zstd{}, nil
		}
	}
	return Stored{}, nil
}

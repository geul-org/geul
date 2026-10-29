//ff:func feature=pack type=writer control=sequence
//ff:what writer.prepareDict — the dictionary new records use: the set's newest, else one trained on this generation when it pays off, else none
package pack

// prepareDict sets w.dict (§6 step 3). Stored and Zstd with DictSize < 0 use
// none. A set that has a dictionary keeps using its newest (compaction drops
// them and trains again). Otherwise one is trained on sampled records when
// there are at least dictMinSamples; when training fails (too little or too
// uniform data) or the dictionary does not save more than its own size, the
// records go without one.
func (w *writer) prepareDict() error {
	z, ok := w.o.Codec.(*Zstd)
	if !ok || z.DictSize < 0 {
		return nil
	}
	if len(w.dicts) > 0 {
		return w.reuseDict(w.dicts[len(w.dicts)-1].id)
	}
	train, test, err := w.samples()
	if err != nil || len(train)+len(test) < dictMinSamples {
		return err
	}
	d, err := z.Train(train, w.nextDictID())
	if err != nil || !w.dictPays(z, d, test) {
		return nil
	}
	return w.writeDict(d)
}

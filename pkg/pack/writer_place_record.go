//ff:func feature=pack type=writer control=sequence
//ff:what writer.placeRecord — append an encoded record to the current data pack (opening a new one when full) with its inventory entry
package pack

// placeRecord stores e as p's record.
func (w *writer) placeRecord(p *put, e encoded) error {
	if w.cur == nil || !w.cur.fits(len(e.b), w.o.PackSize) {
		if err := w.startData(); err != nil {
			return err
		}
	}
	off := w.cur.add(e.b)
	p.loc = Loc{Obj: w.curNum, Off: off, Len: uint64(len(e.b)), Enc: e.enc}
	w.inv.entries = append(w.inv.entries, w.entryOf(p))
	w.res.Records++
	if e.enc == encZstd {
		w.res.Compressed++
	}
	return nil
}

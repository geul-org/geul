//ff:func feature=pack type=writer control=sequence
//ff:what writer.writeRecord — re-read a put, check it did not change, store it in the current data pack (opening a new one when full)
package pack

import (
	"crypto/sha256"
	"fmt"
)

// writeRecord appends p's record and its inventory entry.
func (w *writer) writeRecord(p *put) error {
	data, err := w.b.Get(p.path)
	if err == nil && sha256.Sum256(data) != p.sum {
		err = fmt.Errorf("geul/pack: %s changed while writing", p.path)
	}
	var enc byte
	var stored []byte
	if err == nil {
		enc, stored, err = w.encodeRecord(data)
	}
	if err == nil && (w.cur == nil || !w.cur.fits(len(stored), w.o.PackSize)) {
		err = w.startData()
	}
	if err != nil {
		return err
	}
	off := w.cur.add(stored)
	p.loc = Loc{Obj: w.curNum, Off: off, Len: uint64(len(stored)), Enc: enc}
	w.inv.entries = append(w.inv.entries, w.entryOf(p))
	w.res.Records++
	return nil
}

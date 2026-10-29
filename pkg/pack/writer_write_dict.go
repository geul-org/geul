//ff:func feature=pack type=writer control=sequence
//ff:what writer.writeDict — store a new dictionary raw at the start of the first data pack: inventory kind 4, a DICTS entry
package pack

import "crypto/sha256"

// writeDict opens the first data pack with d and makes it w.dict.
func (w *writer) writeDict(d *Dict) error {
	if err := w.startData(); err != nil {
		return err
	}
	off := w.cur.add(d.Bytes)
	w.inv.entries = append(w.inv.entries, invEntry{kind: invDict, ref: d.ID, n: uint32(len(d.Bytes)), off: off})
	l := Loc{Obj: w.curNum, Off: off, Len: uint64(len(d.Bytes))}
	w.dicts = append(w.dicts, dictEntry{id: d.ID, loc: l, hash: sha256.Sum256(d.Bytes)})
	w.dict, w.res.Dict = d, d.ID
	return nil
}

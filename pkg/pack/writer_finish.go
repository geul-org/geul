//ff:func feature=pack type=writer control=sequence
//ff:what writer.finish — index layers, manifest object, verification on the new generation, then latest.txt
package pack

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// finish writes everything after the data packs (§6 steps 5–7).
func (w *writer) finish() error {
	pl, hl, ns, err := w.layers()
	if err != nil {
		return err
	}
	m := &manifest{gen: w.res.Generation + 1, created: w.o.Now().Unix(), objects: w.objs,
		issuers: w.issuers, named: ns, pLayers: pl, hLayers: hl}
	m.dicts = w.dicts
	if w.prev != nil {
		m.prev, _ = parseHashHex(w.prev.name)
	}
	mb := encodeManifest(m)
	sum := sha256.Sum256(mb)
	name := hex.EncodeToString(sum[:])
	if err := w.writeFile("m/"+name+".gpm", mb); err != nil {
		return err
	}
	if err := w.verify(name); err != nil {
		return err
	}
	if err := w.checkLatest(); err != nil {
		return err
	}
	w.res.Generation, w.res.Manifest, w.res.Changed, w.res.Deleted = m.gen, name, true, len(w.dels)
	w.res.Bytes += int64(len(mb))
	return w.writeFile(latestFile, []byte(fmt.Sprintf("%d %s\n", m.gen, name)))
}

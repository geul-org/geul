//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.nextDictID — a new dictionary ID: one past every ID the set has used, from FirstDictID
package pack

// nextDictID never reuses an ID of the previous generation, even on Fresh.
func (w *writer) nextDictID() uint32 {
	id := uint32(FirstDictID)
	if w.prev == nil {
		return id
	}
	for _, d := range w.prev.m.dicts {
		id = max(id, d.id+1)
	}
	return id
}

//ff:func feature=pack type=writer control=sequence
//ff:what writer.startData — close the open data pack (with its inventory) and open the next
package pack

// startData begins a new data pack with an empty inventory.
func (w *writer) startData() error {
	if w.cur != nil {
		if err := w.finishData(); err != nil {
			return err
		}
	}
	w.inv = inventory{issuers: w.issuers}
	return w.startObject(kindData)
}

//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.writeData — write every record into data packs, then the deletions into the last pack's inventory
package pack

// writeData stores all records and closes the last data pack.
func (w *writer) writeData() error {
	w.assignIssuers()
	for _, p := range w.records() {
		if err := w.writeRecord(p); err != nil {
			return err
		}
	}
	return w.closeData()
}

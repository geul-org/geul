//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.writeData — pick the dictionary, write every record into data packs in batches, then the deletions into the last pack's inventory
package pack

// writeData stores all records and closes the last data pack.
func (w *writer) writeData() error {
	w.assignIssuers()
	if err := w.prepareDict(); err != nil {
		return err
	}
	rs := w.records()
	step := w.workers() * encodeBatch
	for i := 0; i < len(rs); i += step {
		if err := w.writeBatch(rs[i:min(i+step, len(rs))]); err != nil {
			return err
		}
	}
	return w.closeData()
}

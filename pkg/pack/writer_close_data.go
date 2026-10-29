//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.closeData — put the deletions in the last data pack (opening one if no record was written) and close it
package pack

// closeData ends the data phase.
func (w *writer) closeData() error {
	if w.cur == nil {
		if err := w.startData(); err != nil {
			return err
		}
	}
	for _, d := range w.dels {
		w.deleteEntry(d)
	}
	return w.finishData()
}

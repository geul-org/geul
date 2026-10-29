//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what writer.writeBatch — read and encode a run of records in parallel, then place them in order
package pack

// writeBatch stores ps; the bytes written do not depend on the parallelism.
func (w *writer) writeBatch(ps []*put) error {
	es, err := parallelMap(len(ps), w.workers(), func(i int) (encoded, error) { return w.readRecord(ps[i]) })
	if err != nil {
		return err
	}
	for i, p := range ps {
		if err := w.placeRecord(p, es[i]); err != nil {
			return err
		}
	}
	return nil
}

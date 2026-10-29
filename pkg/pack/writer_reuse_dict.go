//ff:func feature=pack type=writer control=sequence
//ff:what writer.reuseDict — use a carried dictionary for new records (its bytes read from the previous generation, SHA-256 checked)
package pack

// reuseDict loads dictionary id of the previous generation into w.dict.
func (w *writer) reuseDict(id uint32) error {
	b, err := w.prev.dict(w.ctx, id)
	if err != nil {
		return err
	}
	w.dict = &Dict{ID: id, Bytes: b}
	return nil
}

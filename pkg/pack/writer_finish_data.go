//ff:func feature=pack type=writer control=sequence
//ff:what writer.finishData — frame the inventory with a checksum and close the data pack
package pack

// finishData writes the inventory (one checksummed zstd frame) and closes the pack.
func (w *writer) finishData() error {
	inv, err := w.o.Codec.Compress(encodeInventory(w.inv))
	if err != nil {
		return err
	}
	return w.closeObject(inv)
}

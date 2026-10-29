//ff:func feature=pack type=writer control=sequence
//ff:what writer.encodeRecord — a record's stored form: zstd (with the generation's dictionary) when that is smaller, raw otherwise (§4.6)
package pack

// encodeRecord picks the encoding; Stored never shrinks bytes, so it is not tried.
func (w *writer) encodeRecord(data []byte) (byte, []byte, error) {
	if _, ok := w.o.Codec.(Stored); ok {
		return encRaw, data, nil
	}
	z, err := w.o.Codec.Compress(data, w.dict)
	if err != nil {
		return 0, nil, err
	}
	if len(z) < len(data) {
		return encZstd, z, nil
	}
	return encRaw, data, nil
}

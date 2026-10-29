//ff:func feature=pack type=helper control=sequence
//ff:what dictBytes — a dictionary's bytes, nil for no dictionary
package pack

// dictBytes returns d.Bytes, or nil when d is nil.
func dictBytes(d *Dict) []byte {
	if d == nil {
		return nil
	}
	return d.Bytes
}

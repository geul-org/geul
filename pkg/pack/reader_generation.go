//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.Generation — the opened manifest's generation number
package pack

// Generation is the manifest's generation (1 for the first).
func (r *Reader) Generation() uint64 {
	return r.m.gen
}

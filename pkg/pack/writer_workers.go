//ff:func feature=pack type=writer control=sequence
//ff:what writer.workers — how many records to read and compress at once: Zstd's workers, 1 otherwise
package pack

// workers is the parallelism of the record phase.
func (w *writer) workers() int {
	if z, ok := w.o.Codec.(*Zstd); ok {
		return z.workers()
	}
	return 1
}

//ff:type feature=pack type=model
//ff:what objWriter — an object being written: temp file, next offset, record or block count, first error
package pack

// objWriter appends to one object; err keeps the first write failure.
type objWriter struct {
	f     objFile
	kind  byte
	off   uint64
	count uint64
	err   error
}

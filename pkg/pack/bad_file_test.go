//ff:type feature=pack type=model
//ff:what badFile — an objFile whose Write, WriteAt, ReadAt or Close fails on demand
package pack

// badFile wraps a real temp file and fails the flagged operations.
type badFile struct {
	objFile
	write, writeAt, readAt, close bool
}

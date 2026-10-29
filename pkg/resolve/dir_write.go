//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.write — store bytes at a path
package resolve

// write stores b at path in the archive.
func (d Dir) write(path string, b []byte) error {
	return d.store().WriteFile(path, b)
}

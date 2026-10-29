//ff:func feature=pack type=writer control=sequence
//ff:what syncFile — fsync a part file; nil is none open
package pack

import "os"

// syncFile flushes f to disk; a nil f (no object open) is nothing to do.
func syncFile(f *os.File) error {
	if f == nil {
		return nil
	}
	return f.Sync()
}

//ff:type feature=pack type=model
//ff:what memFile — an opened Archive file: its info and a seekable reader over its bytes
package pack

import "bytes"

// memFile is the fs.File Open returns for a file (io.Seeker too, for http.ServeFileFS).
type memFile struct {
	info fileInfo
	*bytes.Reader
}

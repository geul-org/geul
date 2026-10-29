//ff:type feature=pack type=model
//ff:what packSet — one pack set of an Archive: its name (Issuer hex, or "" for a set at the top), directory and current generation
package pack

// packSet is a `latest.txt` + `m/` + `o/` directory. r is nil until the set
// has a generation (a set created by a write).
type packSet struct {
	name string
	dir  string
	r    *Reader
}

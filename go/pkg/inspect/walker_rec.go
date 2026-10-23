//ff:func feature=inspect type=decoder control=sequence
//ff:what walker.rec — take one embedded record as a Part, sized by its own Length
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// rec takes the embedded record at the cursor as one Record part. A record
// that does not decode marks the walk bad.
func (w *walker) rec(name string) {
	if w.bad {
		return
	}
	_, n, err := core.DecodeRecord(w.ws[w.pos:])
	if err != nil {
		w.bad = true
		return
	}
	w.out = append(w.out, Part{Name: name, Words: w.ws[w.pos : w.pos+n], Record: true})
	w.pos += n
}

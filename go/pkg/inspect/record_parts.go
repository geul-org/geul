//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what recordParts — the rest of a payload as embedded record parts named by their codecs
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// recordParts takes every remaining record as a Record part named by its spec
// name ("record" when unnamed). The caller has checked the run with recordRun.
func recordParts(w *walker) {
	for w.pos < len(w.ws) && !w.bad {
		r, _, _ := core.DecodeRecord(w.ws[w.pos:])
		name := core.CodecName(r.Kind, r.Codec)
		if name == "" {
			name = "record"
		}
		w.rec(name)
	}
}

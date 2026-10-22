//ff:func feature=geulmd type=encoder control=sequence
//ff:what renderImportDoc — `IMPORT-DOC #id <issuer> <doc id> [v=<version id>]`
package geulmd

import (
	"strconv"

	"github.com/geul-org/geul/go/pkg/core"
)

// renderImportDoc renders Meta IMPORT-DOC: the IMPORT ID at the declared width,
// the Issuer, the 4-word document ID and, when pinned, the 16-word version ID.
// The Issuer width and the version's presence follow from the payload length
// (GEUL t002 OVERVIEW IMPORT).
func renderImportDoc(r core.Record, importWidth int) (string, bool) {
	rest := len(r.Payload) - importWidth - 4
	ver := rest == 17 || rest == 18 || rest == 20 || rest == 32
	iw := rest
	if ver {
		iw -= 16
	}
	if iw != 1 && iw != 2 && iw != 4 && iw != 16 {
		return "", false
	}
	id, _ := core.DecodeID(r.Payload, importWidth)
	p := r.Payload[importWidth:]
	s := "IMPORT-DOC #" + strconv.FormatUint(id, 10) + " " + hexID(p[:iw]) + " " + hexID(p[iw:iw+4])
	if ver {
		s += " v=" + hexID(p[iw+4:])
	}
	return s, true
}

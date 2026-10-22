//ff:func feature=inspect type=decoder control=sequence
//ff:what parseImportDoc — Meta IMPORT-DOC as ImportDoc; false for a layout no Issuer width explains
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// parseImportDoc splits Meta IMPORT-DOC into the IMPORT ID at the declared
// width, the Issuer, the 4-word document ID and, when pinned, the 16-word
// version ID. The Issuer width and the version's presence follow from the
// payload length (GEUL t002 OVERVIEW IMPORT).
func parseImportDoc(r core.Record, importWidth int) (ImportDoc, bool) {
	rest := len(r.Payload) - importWidth - 4
	ver := rest == 17 || rest == 18 || rest == 20 || rest == 32
	iw := rest
	if ver {
		iw -= 16
	}
	if iw != 1 && iw != 2 && iw != 4 && iw != 16 {
		return ImportDoc{}, false
	}
	id, _ := core.DecodeID(r.Payload, importWidth)
	p := r.Payload[importWidth:]
	d := ImportDoc{ID: id, Issuer: p[:iw], Doc: p[iw : iw+4]}
	if ver {
		d.Version = p[iw+4:]
	}
	return d, true
}

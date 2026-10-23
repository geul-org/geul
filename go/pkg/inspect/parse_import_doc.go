//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what parseImportDoc — Meta IMPORT-DOC-* as ImportDoc; false for a payload its codec cannot explain
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// parseImportDoc splits Meta IMPORT-DOC-* into the IMPORT ID, the Issuer whose
// width the codec gives, the 4-word document ID and the 16-word version ID
// (GEUL t002 OVERVIEW IMPORT). The IMPORT ID is the rest, a minimal unsigned
// integer; an all-zero version is "latest" and leaves Version nil.
func parseImportDoc(r core.Record) (ImportDoc, bool) {
	iw := core.ImportIssuerWidth(r.Codec)
	n := len(r.Payload) - iw - 20
	if iw == 0 || n < 1 {
		return ImportDoc{}, false
	}
	id, err := core.DecodeUintValue(core.Record{Kind: core.KindValue, Codec: core.CodecUnsignedInteger, Payload: r.Payload[:n]})
	if err != nil {
		return ImportDoc{}, false
	}
	p := r.Payload[n:]
	d := ImportDoc{ID: id, Issuer: p[:iw], Doc: p[iw : iw+4]}
	for _, v := range p[iw+4:] {
		if v != 0 {
			d.Version = p[iw+4:]
		}
	}
	return d, true
}

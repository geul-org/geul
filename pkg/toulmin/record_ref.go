//ff:func feature=toulmin type=decoder control=sequence
//ff:what RecordRef — the TANGEUL CodecRef of a decoded Edge record (GEUL ISSUER formats only)
package toulmin

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// RecordRef returns the CodecRef of r. TANGEUL records are Edge records; an
// ISSUER record is accepted only under the 1-word GEUL Issuer (0x0000), since
// every TANGEUL format number is defined there.
func RecordRef(r core.Record) (CodecRef, error) {
	if r.Kind != core.KindEdge {
		return CodecRef{}, fmt.Errorf("geul/toulmin: expected an Edge record, got kind %d", r.Kind)
	}
	if core.IssuerWidth(r.Kind, r.Codec) == 0 {
		return CodecRef{Codec: r.Codec}, nil
	}
	if r.Codec != core.CodecIssuer1Word || uint16(r.Issuer[0]) != IssuerGEUL {
		return CodecRef{}, fmt.Errorf("geul/toulmin: ISSUER record 0x%02X is not under the GEUL Issuer", r.Codec)
	}
	return CodecRef{Codec: r.Codec, Format: r.Format}, nil
}

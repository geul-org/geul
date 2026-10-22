//ff:func feature=geulmd type=encoder control=sequence
//ff:what rawRecord — the lossless fallback line of any record: RECORD kind codec [issuer format] : payload
package geulmd

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// rawRecord renders r without interpreting its payload: `RECORD <kind> 0x<codec>`,
// the Issuer and format number when the codec is an ISSUER code, then `:` and
// the payload words. Length is recomputed on encoding, so it is not written.
func rawRecord(r core.Record) string {
	s := fmt.Sprintf("RECORD %s 0x%02x", kindName(r.Kind), r.Codec)
	if len(r.Issuer) > 0 {
		s += fmt.Sprintf(" %s 0x%04x", hexID(r.Issuer), r.Format)
	}
	if len(r.Payload) == 0 {
		return s + " :"
	}
	return s + " : " + hexWords(r.Payload)
}

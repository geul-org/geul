//ff:func feature=schema type=helper control=sequence
//ff:what subjectOf — test fixture: EXTERNAL-ISSUER-1WORD Reference to Issuer 0x0000, a document and a target
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
)

func subjectOf(doc, target core.Word) core.Record {
	return core.Record{Kind: core.KindReference, Codec: core.RefExternalIssuer1Word << 3, Payload: []core.Word{0, 0xffff, 0, 0, doc, target}}
}

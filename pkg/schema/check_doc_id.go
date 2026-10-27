//ff:func feature=schema type=rule control=sequence
//ff:what checkDocID — M2: DOC-ID is one EXTERNAL-ISSUER-*(Issuer, document ID, 0) Reference
package schema

import "github.com/geul-org/gogeul/pkg/core"

// checkDocID applies M2 to the payload of a DOC-ID.
func checkDocID(r rec) []Fact {
	if _, _, err := core.DecodeDocID(r.R); err != nil {
		return []Fact{factAt("M2", Error, r, "Payload", "DOC-POINTER(Issuer, 문서 ID, 0) Reference 하나", err.Error())}
	}
	return nil
}

//ff:func feature=core type=helper control=sequence
//ff:what IssuerFormatName — the spec name of a Value ISSUER format (1-WORD Issuer and format number); "" when unknown
package core

// IssuerFormatName returns the spec name of the ISSUER format of a Value
// record r (grammar Value ISSUER 형식): its 1-WORD Issuer and format number
// looked up in valueIssuerFormatNames. Other kinds, wider Issuers and
// unassigned formats return "".
func IssuerFormatName(r Record) string {
	if r.Kind != KindValue || len(r.Issuer) != 1 {
		return ""
	}
	return valueIssuerFormatNames[uint32(r.Issuer[0])<<16|uint32(r.Format)]
}

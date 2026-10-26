//ff:func feature=core type=helper control=sequence
//ff:what IsAbsent — whether a decoded record is the Absent record (TAG WORD 0x0000, no payload)
package core

// IsAbsent reports whether r is the Absent record: Reference Codec 0000,
// Quantifier 000, no payload and no ISSUER extension.
func IsAbsent(r Record) bool {
	return r.Kind == KindReference && r.Codec == 0 && len(r.Payload) == 0 && len(r.Issuer) == 0
}

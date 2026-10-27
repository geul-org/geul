//ff:func feature=core type=helper control=sequence
//ff:what RecordName — the spec name of a record: its Value ISSUER format name when known, else its codec name
package core

// RecordName names r for display: a Value ISSUER format by its format name
// (IssuerFormatName, e.g. RATIONAL-I64), any other record by CodecName.
func RecordName(r Record) string {
	if name := IssuerFormatName(r); name != "" {
		return name
	}
	return CodecName(r.Kind, r.Codec)
}

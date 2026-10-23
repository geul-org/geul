//ff:func feature=schema type=rule control=sequence
//ff:what checkDocName — M2: DOC-NAME is a 1/2/4/16-word Issuer and a 4-word document ID
package schema

// checkDocName applies M2 to the payload of a DOC-NAME.
func checkDocName(r rec) []Fact {
	w := len(r.R.Payload) - 4
	if w != 1 && w != 2 && w != 4 && w != 16 {
		return []Fact{factAt("M2", Error, r, "Payload", "Issuer 1/2/4/16 WORD + 문서 ID 4 WORD", fmtWords(len(r.R.Payload)))}
	}
	return nil
}

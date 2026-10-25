//ff:func feature=schema type=rule control=sequence
//ff:what checkImportDoc — M3: IMPORT-DOC is a DOC-POINTER and a minimal IMPORT ID
package schema

// checkImportDoc applies M3 to one IMPORT-DOC: a DOC-POINTER Reference of the
// imported document, then the IMPORT ID (minimal, 1 WORD or more).
func checkImportDoc(r rec) []Fact {
	ws, ok := importSplit(r)
	if !ok {
		return []Fact{factAt("M3", Error, r, "Payload", "DOC-POINTER(Issuer, 문서 ID, 버전 ID) + IMPORT ID", fmtWords(len(r.R.Payload)))}
	}
	if !minimalUnsigned(ws) {
		return []Fact{factAt("M3", Error, r, "IMPORT ID", "1 WORD 이상, 최소 WORD", fmtWords(len(ws)))}
	}
	return nil
}

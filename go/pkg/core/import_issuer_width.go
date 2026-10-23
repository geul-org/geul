//ff:func feature=core type=helper control=selection
//ff:what ImportIssuerWidth — the Issuer word count of an IMPORT-DOC-* Meta codec, 0 otherwise
package core

// ImportIssuerWidth returns the Issuer width that an IMPORT-DOC-1/2/4/16WORD
// Meta codec declares (GEUL t002 OVERVIEW IMPORT): its low 2 bits, 00=1, 01=2,
// 10=4, 11=16 words. Any other codec returns 0.
func ImportIssuerWidth(codec uint8) int {
	switch codec {
	case MetaImportDoc1Word:
		return 1
	case MetaImportDoc2Word:
		return 2
	case MetaImportDoc4Word:
		return 4
	case MetaImportDoc16Word:
		return 16
	default:
		return 0
	}
}

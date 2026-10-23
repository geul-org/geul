//ff:func feature=core type=helper control=selection
//ff:what CodecName — the spec name of a Reference, Meta, Value, execution Edge or ISSUER codec; "" otherwise
package core

// metaNames maps the Meta codecs this package defines to their spec names.
var metaNames = map[uint8]string{
	MetaDoc1Word: "DOC-1WORD", MetaDoc2Word: "DOC-2WORD", MetaDoc4Word: "DOC-4WORD",
	MetaBody: "BODY", MetaEnd: "END", MetaSpecVersion: "SPEC-VERSION", MetaProfile: "PROFILE",
	MetaDocName: "DOC-NAME", MetaDefaultLanguage: "DEFAULT-LANGUAGE",
	MetaImportDoc1Word: "IMPORT-DOC-1WORD", MetaImportDoc2Word: "IMPORT-DOC-2WORD",
	MetaImportDoc4Word: "IMPORT-DOC-4WORD", MetaImportDoc16Word: "IMPORT-DOC-16WORD",
}

// refNames maps the 4-bit Reference codecs (CCCC) to their spec names.
var refNames = map[uint8]string{
	RefExternalPointer1Word: "EXTERNAL-POINTER-1WORD", RefExternalPointer2Word: "EXTERNAL-POINTER-2WORD",
	RefExternalPointer4Word: "EXTERNAL-POINTER-4WORD", RefExternalPointer8Word: "EXTERNAL-POINTER-8WORD",
	RefExternalIssuer1Word: "EXTERNAL-ISSUER-1WORD", RefExternalIssuer2Word: "EXTERNAL-ISSUER-2WORD",
	RefExternalIssuer4Word: "EXTERNAL-ISSUER-4WORD", RefExternalIssuer16Word: "EXTERNAL-ISSUER-16WORD",
	RefLocal: "LOCAL", RefGeul: "GEUL", RefWikidataQID: "WIKIDATA-QID", RefWikidataPID: "WIKIDATA-PID",
}

// edgeNames maps the execution-family direct Edge codecs (GEUL t002 OVERVIEW
// 실행) to their spec names. Profile families such as TOULMIN name their own.
var edgeNames = map[uint8]string{
	0x40: "SYMBOL-DECLARE", 0x41: "FUNCTION-DECLARE", 0x42: "TYPE-DECLARE",
	0x48: "CONST", 0x49: "UNARY", 0x4A: "BINARY", 0x4B: "CALL", 0x4C: "MEMBER", 0x4D: "INDEX",
	0x4E: "CONSTRUCT", 0x4F: "CONDITIONAL", 0x50: "BLOCK", 0x51: "ASSIGN", 0x52: "IF", 0x53: "SWITCH",
	0x54: "WHILE", 0x55: "FOREACH", 0x56: "JUMP", 0x57: "RETURN", 0x58: "RAISE", 0x59: "TRY",
}

// valueNames maps the direct Value codecs this package defines to their spec
// names. The odd member of an even/odd byte-string pair shares the name.
var valueNames = map[uint8]string{
	CodecUnsignedInteger: "UNSIGNED-INTEGER", CodecSignedDecimal: "SIGNED-DECIMAL",
	CodecUnsignedDecimal: "UNSIGNED-DECIMAL", CodecUTF8TextEven: "UTF8-TEXT", CodecUTF8TextEven | 1: "UTF8-TEXT",
	CodecLanguageBCP47Even: "LANGUAGE-BCP47-CODE", CodecLanguageBCP47Even | 1: "LANGUAGE-BCP47-CODE",
}

// issuerNames maps the ISSUER codecs of the Value and Edge kinds to their names.
var issuerNames = map[uint8]string{
	CodecIssuer1Word: "ISSUER-1WORD", CodecIssuer2Word: "ISSUER-2WORD",
	CodecIssuer4Word: "ISSUER-4WORD", CodecIssuer16Word: "ISSUER-16WORD",
}

// CodecName returns the spec name of codec under kind k (GEUL t002 OVERVIEW):
// a Reference code (named by CCCC, ignoring the Quantifier), a Meta code, a
// direct Value code, an execution Edge code or an ISSUER code. Codes defined by
// a profile (such as the TOULMIN Edge codes) and unknown codes return "".
func CodecName(k Kind, codec uint8) string {
	switch {
	case k == KindReference:
		return refNames[codec>>3]
	case k == KindMeta:
		return metaNames[codec]
	case IssuerWidth(k, codec) > 0:
		return issuerNames[codec]
	case k == KindValue:
		return valueNames[codec]
	default:
		return edgeNames[codec]
	}
}

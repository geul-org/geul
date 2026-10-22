//ff:func feature=core type=helper control=selection
//ff:what CodecName — the spec name of a Meta, Value or ISSUER codec this package knows; "" otherwise
package core

// metaNames maps the Meta codecs this package defines to their spec names.
var metaNames = map[uint8]string{
	MetaDoc1Word: "DOC-1WORD", MetaDoc2Word: "DOC-2WORD", MetaDoc4Word: "DOC-4WORD",
	MetaBody: "BODY", MetaEnd: "END", MetaSpecVersion: "SPEC-VERSION", MetaProfile: "PROFILE",
	MetaDocName: "DOC-NAME", MetaIDWidth: "ID-WIDTH", MetaDefaultLanguage: "DEFAULT-LANGUAGE",
	MetaImportDoc: "IMPORT-DOC",
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
// a Meta code, a direct Value code or an ISSUER code. Codes defined by a
// profile (such as the TOULMIN Edge codes) and unknown codes return "".
func CodecName(k Kind, codec uint8) string {
	switch {
	case k == KindMeta:
		return metaNames[codec]
	case IssuerWidth(k, codec) > 0:
		return issuerNames[codec]
	case k == KindValue:
		return valueNames[codec]
	default:
		return ""
	}
}

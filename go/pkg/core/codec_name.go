//ff:func feature=core type=helper control=selection
//ff:what CodecName — the spec name of a Reference, Meta, Value, OVERVIEW Edge or ISSUER codec; "" otherwise
package core

// metaNames maps the Meta codecs (GEUL t002 OVERVIEW Meta, docs/meta) to their spec names.
var metaNames = map[uint8]string{
	MetaDoc1Word: "DOC-1WORD", MetaDoc2Word: "DOC-2WORD", MetaDoc4Word: "DOC-4WORD",
	MetaBody: "BODY", MetaEnd: "END", MetaSpecVersion: "SPEC-VERSION", MetaProfile: "PROFILE",
	MetaDocID: "DOC-ID", MetaEffectiveFrom: "EFFECTIVE-FROM", MetaEffectiveUntil: "EFFECTIVE-UNTIL", MetaDefaultLanguage: "DEFAULT-LANGUAGE",
	MetaImportDoc: "IMPORT-DOC", 0x20: "AUTHOR", 0x21: "CO-AUTHOR", 0x22: "CREATED-AT", 0x23: "MODIFIED-AT", 0x24: "RECORDED-BY",
	0x25: "RECORDED-AT", 0x26: "SOURCE", 0x27: "SOURCE-SPAN", 0x28: "ACQUIRED-AT",
	0x30: "PARENT", 0x31: "DERIVED-FROM", 0x32: "GENERATED-BY", 0x33: "TRANSFORMED-BY", 0x34: "CHERRY-PICKED-FROM",
	0x35: "REVERTS", 0x36: "REPLACES", 0x37: "SQUASHED-FROM", 0x38: "MESSAGE", 0x39: "TAG", MetaDeltaBase: "DELTA-BASE", MetaDeltaRemove: "DELTA-REMOVE", MetaNextLocalID: "NEXT-LOCAL-ID",
	MetaIssuerKey: "ISSUER-KEY", 0x41: "SIGNATURE", 0x42: "SIGNED-OFF", 0x43: "CONFIDENCE", 0x44: "VERDICT",
	0x50: "LICENSE", 0x51: "ACCESS", 0x52: "PRIVACY", 0x53: "RETENTION", 0x54: "CONSENT", 0x55: "USAGE",
}

// refNames maps the 4-bit Reference codecs (CCCC) to their spec names.
var refNames = map[uint8]string{
	RefExternalPointer1Word: "EXTERNAL-POINTER-1WORD", RefExternalPointer2Word: "EXTERNAL-POINTER-2WORD",
	RefExternalPointer4Word: "EXTERNAL-POINTER-4WORD", RefExternalPointer8Word: "EXTERNAL-POINTER-8WORD",
	RefExternalIssuer1Word: "EXTERNAL-ISSUER-1WORD", RefExternalIssuer2Word: "EXTERNAL-ISSUER-2WORD",
	RefExternalIssuer4Word: "EXTERNAL-ISSUER-4WORD", RefExternalIssuer16Word: "EXTERNAL-ISSUER-16WORD",
	RefGeul: "GEUL", RefGeulSIDX: "GEUL-SIDX", RefWikidataQID: "WIKIDATA-QID", RefWikidataPID: "WIKIDATA-PID", RefLocal: "LOCAL",
	RefIssuerPointer: "ISSUER-POINTER", RefDocPointer: "DOC-POINTER",
}

// edgeNames maps the direct Edge codecs of the 서술, 근거 and 실행 families
// (GEUL t002 OVERVIEW Edge) to their spec names. Profile families such as
// TOULMIN name their own.
var edgeNames = map[uint8]string{
	0x00: "REFERENCE-NODE", 0x01: "SAME-NODE", 0x02: "DIFFERENT-NODE", EdgeBind: "BIND", 0x04: "LABEL", 0x05: "ALIAS", 0x06: "DESCRIPTION",
	0x08: "VERB-SHORT", 0x09: "VERB", 0x0A: "TRIPLE", 0x0B: "QUALIFIER", 0x10: "WORLD", 0x11: "TIMELINE", 0x12: "SCENE",
	EdgePublication: "PUBLICATION", EdgeDeprecation: "DEPRECATION", EdgeTombstone: "TOMBSTONE",
	EdgeDeclare: "DECLARE", EdgeField: "FIELD", EdgeEnum: "ENUM", EdgeCodec: "CODEC",
	0x20: "RECORD-SOURCE-SPAN", 0x21: "RECORD-SOURCE", 0x22: "RECORD-DERIVED-FROM", 0x23: "RECORD-GENERATED-BY",
	0x24: "RECORD-INTRODUCED-IN", 0x25: "RECORD-CONFIDENCE", 0x26: "RECORD-EFFECTIVE",
	0x41: "FUNCTION-DECLARE", 0x42: "TYPE-DECLARE",
	0x48: "CONST", 0x49: "UNARY", 0x4A: "BINARY", 0x4B: "CALL", 0x4C: "MEMBER", 0x4D: "INDEX",
	0x4E: "CONSTRUCT", 0x4F: "CONDITIONAL", 0x50: "BLOCK", 0x51: "ASSIGN", 0x52: "IF", 0x53: "SWITCH",
	0x54: "WHILE", 0x55: "FOREACH", 0x56: "JUMP", 0x57: "RETURN", 0x58: "RAISE", 0x59: "TRY",
}

// valueNames maps the direct Value codecs (GEUL t002 docs/value) to their spec
// names. The odd member of an even/odd byte-string pair shares the name.
var valueNames = map[uint8]string{
	0x00: "BOOL", 0x01: "SIGNED-INTEGER", CodecUnsignedInteger: "UNSIGNED-INTEGER", 0x03: "BINARY-FLOAT",
	CodecSignedDecimal: "SIGNED-DECIMAL", CodecUnsignedDecimal: "UNSIGNED-DECIMAL",
	CodecGregorianDate: "GREGORIAN-DATE", 0x21: "CIVIL-TIME-NANOSECONDS", CodecPosixSeconds: "POSIX-TIMESTAMP-SECONDS", 0x23: "DURATION-SECONDS",
	CodecEDTFEven: "EDTF", CodecEDTFEven | 1: "EDTF",
	CodecGEDCOMDateEven: "GEDCOM-DATE", CodecGEDCOMDateEven | 1: "GEDCOM-DATE",
	CodecUTF8TextEven: "UTF8-TEXT", CodecUTF8TextEven | 1: "UTF8-TEXT",
	CodecCoordinateEarth: "COORDINATE-EARTH", CodecCoordinateGeo: "COORDINATE-GEO", CodecQuantity: "QUANTITY",
	CodecUTF8LanguageTextEven: "UTF8-LANGUAGE-TEXT", CodecUTF8LanguageTextEven | 1: "UTF8-LANGUAGE-TEXT",
	0x44: "URI", 0x45: "URI", CodecLanguageBCP47Even: "LANGUAGE-BCP47-CODE", CodecLanguageBCP47Even | 1: "LANGUAGE-BCP47-CODE",
	0x48: "LANGUAGE-SHORT-CODE", 0x60: "UUID", CodecSHA256Digest: "SHA-256-DIGEST", CodecSHA1Digest: "SHA1-DIGEST",
}

// issuerNames maps the ISSUER codecs of the Value and Edge kinds to their names.
var issuerNames = map[uint8]string{
	CodecIssuer1Word: "ISSUER-1WORD", CodecIssuer2Word: "ISSUER-2WORD",
	CodecIssuer4Word: "ISSUER-4WORD", CodecIssuer16Word: "ISSUER-16WORD",
}

// CodecName returns the spec name of codec under kind k (GEUL t002 OVERVIEW):
// a Reference code (named by CCCC, ignoring the Quantifier), a Meta code, a
// direct Value code, a 서술/근거/실행 Edge code or an ISSUER code. Codes defined by
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

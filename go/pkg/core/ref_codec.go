//ff:type feature=core type=model
//ff:what Reference codec constants — the GEUL t002 4-bit Reference codes (CCCC) and the LOCAL TAG codec
package core

// Reference codecs (GEUL t002 OVERVIEW Reference). A Reference TAG WORD holds
// the 4-bit codec CCCC and the 3-bit Quantifier QQQ in its 7 codec bits, so the
// Record codec is CCCC<<3 | QQQ. 0000 is reserved, so no Reference TAG is the
// WORD 0x0000. The top bit splits the codes: 0xxx identify their target
// directly — an ID this document issues (LOCAL), an issued ID (GEUL,
// GEUL-SIDX, WIKIDATA-*), or an Issuer or one of its documents
// (ISSUER-POINTER, DOC-POINTER, which carry no target ID) — and 1xxx point at
// another document's LOCAL ID: 10WW with an IMPORT ID of 1/2/4/8 words, 11WW
// with an Issuer of 1/2/4/16 words and a 4-word document ID.
const (
	// RefLocal is an ID this document issues (Edge Ids included).
	RefLocal uint8 = 0x1
	// RefGeul is an entity geul.org issues under CC0.
	RefGeul uint8 = 0x2
	// RefGeulSIDX is a GEUL Semantic-aligned Index (SIDX).
	RefGeulSIDX uint8 = 0x3
	// RefWikidataQID is a Wikidata item (the integer after Q).
	RefWikidataQID uint8 = 0x4
	// RefWikidataPID is a Wikidata property (the integer after P).
	RefWikidataPID uint8 = 0x5
	// RefIssuerPointer is an Issuer itself; its width is the Length minus the TAG.
	RefIssuerPointer uint8 = 0x6
	// RefDocPointer is one version of an Issuer's document: Issuer, 4-word
	// document ID and 2-word version ID (0 is the latest).
	RefDocPointer uint8 = 0x7
	// RefExternalPointer1Word points at an imported document's LOCAL ID (1-word IMPORT ID).
	RefExternalPointer1Word uint8 = 0x8
	// RefExternalPointer2Word points with a 2-word IMPORT ID.
	RefExternalPointer2Word uint8 = 0x9
	// RefExternalPointer4Word points with a 4-word IMPORT ID.
	RefExternalPointer4Word uint8 = 0xA
	// RefExternalPointer8Word points with an 8-word IMPORT ID.
	RefExternalPointer8Word uint8 = 0xB
	// RefExternalIssuer1Word cites a 1-word Issuer document's LOCAL ID.
	RefExternalIssuer1Word uint8 = 0xC
	// RefExternalIssuer2Word cites a 2-word Issuer document's LOCAL ID.
	RefExternalIssuer2Word uint8 = 0xD
	// RefExternalIssuer4Word cites a 4-word Issuer document's LOCAL ID.
	RefExternalIssuer4Word uint8 = 0xE
	// RefExternalIssuer16Word cites a 16-word Issuer document's LOCAL ID.
	RefExternalIssuer16Word uint8 = 0xF
	// CodecLocalIdentified is the Record codec of a LOCAL Reference with
	// Quantifier Identified (QQQ=000): the form every Edge Id declaration uses.
	CodecLocalIdentified uint8 = RefLocal << 3
)

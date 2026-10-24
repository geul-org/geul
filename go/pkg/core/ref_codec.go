//ff:type feature=core type=model
//ff:what Reference codec constants — the GEUL t002 4-bit Reference codes (CCCC) and the LOCAL TAG codec
package core

// Reference codecs (GEUL t002 OVERVIEW Reference). A Reference TAG WORD holds
// the 4-bit codec CCCC and the 3-bit Quantifier QQQ in its 7 codec bits, so the
// Record codec is CCCC<<3 | QQQ. 지시 codes (00xx) carry an IMPORT ID of
// 1/2/4/8 words, 지시선언 codes (01xx) an Issuer of 1/2/4/16 words plus a
// 4-word document ID, and 선언 codes (1000-1100) only the target ID. The
// 지시선언 codes ISSUER-POINTER (1101) and DOC-POINTER (1110) carry no target
// ID: an Issuer, or an Issuer, a 4-word document ID and a 2-word version ID.
const (
	// RefExternalPointer1Word points at an imported document's LOCAL ID (1-word IMPORT ID).
	RefExternalPointer1Word uint8 = 0x0
	// RefExternalPointer2Word points with a 2-word IMPORT ID.
	RefExternalPointer2Word uint8 = 0x1
	// RefExternalPointer4Word points with a 4-word IMPORT ID.
	RefExternalPointer4Word uint8 = 0x2
	// RefExternalPointer8Word points with an 8-word IMPORT ID.
	RefExternalPointer8Word uint8 = 0x3
	// RefExternalIssuer1Word cites a 1-word Issuer document's LOCAL ID.
	RefExternalIssuer1Word uint8 = 0x4
	// RefExternalIssuer2Word cites a 2-word Issuer document's LOCAL ID.
	RefExternalIssuer2Word uint8 = 0x5
	// RefExternalIssuer4Word cites a 4-word Issuer document's LOCAL ID.
	RefExternalIssuer4Word uint8 = 0x6
	// RefExternalIssuer16Word cites a 16-word Issuer document's LOCAL ID.
	RefExternalIssuer16Word uint8 = 0x7
	// RefGeul is an entity geul.org issues under CC0.
	RefGeul uint8 = 0x8
	// RefGeulSIDX is a GEUL Semantic-aligned Index (SIDX).
	RefGeulSIDX uint8 = 0x9
	// RefWikidataQID is a Wikidata item (the integer after Q).
	RefWikidataQID uint8 = 0xA
	// RefWikidataPID is a Wikidata property (the integer after P).
	RefWikidataPID uint8 = 0xB
	// RefLocal is an ID this document issues (Edge Ids included).
	RefLocal uint8 = 0xC
	// RefIssuerPointer is an Issuer itself; its width is the Length minus the TAG.
	RefIssuerPointer uint8 = 0xD
	// RefDocPointer is one version of an Issuer's document: Issuer, 4-word
	// document ID and 2-word version ID (0 is the latest).
	RefDocPointer uint8 = 0xE
	// CodecLocalIdentified is the Record codec of a LOCAL Reference with
	// Quantifier Identified (QQQ=000): the form every Edge Id declaration uses.
	CodecLocalIdentified uint8 = RefLocal << 3
)

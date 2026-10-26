//ff:type feature=core type=model
//ff:what Meta codec constants — the GEUL t002 structure, declaration and diff Meta codes the library uses
package core

// Meta codecs from the GEUL t002 OVERVIEW Meta tables used by a TANGEUL
// document: the structure class (document boundary) and the declaration class.
const (
	// MetaDoc1Word opens a document whose total length fits 1 word.
	MetaDoc1Word uint8 = 0x00
	// MetaDoc2Word opens a document whose total length fits 2 words.
	MetaDoc2Word uint8 = 0x01
	// MetaDoc4Word opens a document whose total length fits 4 words.
	MetaDoc4Word uint8 = 0x02
	// MetaBody ends the HEAD and starts the BODY (1 word).
	MetaBody uint8 = 0x0E
	// MetaEnd closes the document (1 word).
	MetaEnd uint8 = 0x0F
	// MetaSpecVersion declares the GEUL spec version followed.
	MetaSpecVersion uint8 = 0x10
	// MetaProfile declares the applied profile.
	MetaProfile uint8 = 0x11
	// MetaDocID identifies this document: Issuer + 4-word document ID (formerly DOC-NAME).
	MetaDocID uint8 = 0x12
	// MetaEffectiveFrom declares when this version takes effect (시행일).
	MetaEffectiveFrom uint8 = 0x13
	// MetaEffectiveUntil declares when this version stops being in effect (한시법).
	MetaEffectiveUntil uint8 = 0x14
	// MetaDefaultLanguage declares the default language.
	MetaDefaultLanguage uint8 = 0x16
	// MetaImportDoc imports a document under an IMPORT ID: a DOC-POINTER
	// Reference, then the IMPORT ID.
	MetaImportDoc uint8 = 0x60
	// MetaDeltaBase marks a version file as a diff against an earlier version
	// of the same document: a 2-word version ID.
	MetaDeltaBase uint8 = 0x3A
	// MetaDeltaRemove lists the LOCAL IDs a diff removes from its base.
	MetaDeltaRemove uint8 = 0x3B
	// MetaNextLocalID is the next LOCAL ID a document assigns (auto
	// increment), kept in its version list.
	MetaNextLocalID uint8 = 0x3C
)

//ff:func feature=inspect type=decoder control=selection
//ff:what headEntry — one HEAD entry with its typed content for the Meta this library knows
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// headEntry interprets a HEAD Meta record. SPEC-VERSION is a bare unsigned
// integer, shown as an UNSIGNED-INTEGER Value; PROFILE and DEFAULT-LANGUAGE
// carry one Value; DOC-ID and IMPORT-DOC-* their ids. Any
// other Meta, or a known one whose payload does not fit its form, keeps Data nil.
func headEntry(r core.Record) Entry {
	e := newEntry(r)
	switch r.Codec {
	case core.MetaSpecVersion:
		e.Data = asData(specVersionValue(r))
	case core.MetaProfile, core.MetaDefaultLanguage:
		e.Data = asData(metaValue(r))
	case core.MetaDocID:
		e.Data = asData(parseDocID(r))
	case core.MetaImportDoc1Word, core.MetaImportDoc2Word, core.MetaImportDoc4Word, core.MetaImportDoc16Word:
		e.Data = asData(parseImportDoc(r))
	}
	return e
}

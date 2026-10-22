//ff:func feature=inspect type=decoder control=selection
//ff:what headEntry — one HEAD entry with its typed content for the Meta this library knows
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// headEntry interprets a HEAD Meta record. SPEC-VERSION, PROFILE and
// DEFAULT-LANGUAGE carry one Value, ID-WIDTH its Widths, DOC-NAME and
// IMPORT-DOC their ids. Any other Meta, or a known one whose payload does not
// fit its form, keeps Data nil.
func headEntry(r core.Record, w Widths) Entry {
	e := newEntry(r)
	switch r.Codec {
	case core.MetaSpecVersion, core.MetaProfile, core.MetaDefaultLanguage:
		e.Data = asData(metaValue(r))
	case core.MetaIDWidth:
		e.Data = asData(parseWidths(r))
	case core.MetaDocName:
		e.Data = asData(parseDocName(r))
	case core.MetaImportDoc:
		e.Data = asData(parseImportDoc(r, w.Import))
	}
	return e
}

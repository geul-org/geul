//ff:func feature=inspect type=decoder control=selection
//ff:what headEntry — one HEAD entry with its typed content for the Meta this library knows; a HEAD Edge stays bare
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// headEntry interprets a HEAD Meta record. SPEC-VERSION is a bare unsigned
// integer, shown as an UNSIGNED-INTEGER Value; DOC-ID and IMPORT-DOC carry
// their ids; any other Meta whose payload is exactly one known Value (PROFILE,
// PARENT, RECORDED-AT, EFFECTIVE-FROM …) carries that Value. The rest, or a
// known one whose payload does not fit its form, keeps Data nil. A HEAD Edge
// (a SOURCE on the whole document or a QUALIFIER on one) is bare like a BODY
// Edge; Decode gives it its Edge Id.
func headEntry(r core.Record) Entry {
	e := newEntry(r)
	switch {
	case r.Kind != core.KindMeta:
	case r.Codec == core.MetaSpecVersion:
		e.Data = asData(specVersionValue(r))
	case r.Codec == core.MetaDocID:
		e.Data = asData(parseDocID(r))
	case r.Codec == core.MetaImportDoc:
		e.Data = asData(parseImportDoc(r))
	default:
		e.Data = asData(metaValue(r))
	}
	return e
}

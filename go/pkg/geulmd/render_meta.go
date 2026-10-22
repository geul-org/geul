//ff:func feature=geulmd type=encoder control=selection
//ff:what renderMeta — one HEAD line: a typed form for the Meta GEULMD knows, a raw line otherwise
package geulmd

import (
	"strings"

	"github.com/geul-org/geul/go/pkg/core"
)

// renderMeta renders a HEAD Meta record. SPEC-VERSION, PROFILE and
// DEFAULT-LANGUAGE show their Value, ID-WIDTH its widths, DOC-NAME and
// IMPORT-DOC their ids; every other Meta, or a known one whose payload does not
// fit its form, becomes a raw RECORD line. PROFILE "tangeul-…" turns on
// structured TOULMIN rendering of the BODY.
func renderMeta(r core.Record, ctx *headCtx) string {
	var s string
	ok := false
	switch r.Codec {
	case core.MetaSpecVersion:
		s, ok = metaValue("SPEC-VERSION", r)
	case core.MetaProfile:
		s, ok = metaValue("PROFILE", r)
		ctx.tangeul = ok && strings.HasPrefix(s, `PROFILE "tangeul-`)
	case core.MetaDefaultLanguage:
		s, ok = metaValue("DEFAULT-LANGUAGE", r)
	case core.MetaIDWidth:
		s, ok = renderIDWidth(r, ctx)
	case core.MetaDocName:
		s, ok = renderDocName(r)
	case core.MetaImportDoc:
		s, ok = renderImportDoc(r, ctx.importWidth)
	}
	if !ok {
		return rawRecord(r)
	}
	return s
}

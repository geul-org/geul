//ff:func feature=geulmd type=encoder control=selection
//ff:what renderMeta — one HEAD line: a typed form for interpreted Meta, a raw line otherwise
package geulmd

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

// renderMeta renders a HEAD entry. SPEC-VERSION, PROFILE and DEFAULT-LANGUAGE
// show their Value, ID-WIDTH its widths, DOC-NAME and IMPORT-DOC their ids;
// every other Meta, or a known one whose payload did not fit its form, becomes
// a raw RECORD line.
func renderMeta(e inspect.Entry) string {
	switch v := e.Data.(type) {
	case inspect.Value:
		return e.Name + " " + renderValue(v)
	case inspect.Widths:
		return fmt.Sprintf("ID-WIDTH edge=%d var=%d import=%d", v.Edge, v.Var, v.Import)
	case inspect.DocName:
		return "DOC-NAME " + core.FormatID(v.Issuer) + " " + core.FormatID(v.Doc)
	case inspect.ImportDoc:
		return renderImportDoc(v)
	default:
		return rawRecord(e.Record)
	}
}

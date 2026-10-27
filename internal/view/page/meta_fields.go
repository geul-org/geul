//ff:func feature=page type=builder control=selection
//ff:what metaFields — the fields of a HEAD Meta record from its interpreted content
package page

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// metaFields lays out an interpreted Meta: a Value, a DOC-ID or an
// IMPORT-DOC with its pinned version, or "latest" when the version is 0. A Meta with no typed reading shows its codec spec fields, else its Issuer and format, if any.
func metaFields(e inspect.Entry) []Field {
	switch v := e.Data.(type) {
	case inspect.Value:
		return []Field{{Key: "value", Text: valueText(v)}}
	case inspect.DocID:
		return []Field{{Key: "issuer", Text: core.FormatID(v.Issuer)}, {Key: "doc", Text: core.FormatID(v.Doc)}}
	case inspect.ImportDoc:
		ver := "latest"
		if v.Version != 0 {
			ver = core.FormatID([]core.Word{core.Word(v.Version >> 16), core.Word(v.Version)})
		}
		return []Field{{Key: "import", Text: "#" + strconv.FormatUint(v.ID, 10)}, {Key: "issuer", Text: core.FormatID(v.Issuer)}, {Key: "doc", Text: core.FormatID(v.Doc)}, {Key: "version", Text: ver}}
	default:
		return specFields(e)
	}
}

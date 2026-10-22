//ff:func feature=geulmd type=encoder control=sequence
//ff:what renderImportDoc — IMPORT-DOC as `#id issuer doc [v=version]`
package geulmd

import (
	"strconv"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

// renderImportDoc renders Meta IMPORT-DOC: the IMPORT ID, the Issuer, the
// 4-word document ID and, when pinned, the 16-word version ID.
func renderImportDoc(v inspect.ImportDoc) string {
	s := "IMPORT-DOC #" + strconv.FormatUint(v.ID, 10) + " " + core.FormatID(v.Issuer) + " " + core.FormatID(v.Doc)
	if v.Version != nil {
		s += " v=" + core.FormatID(v.Version)
	}
	return s
}

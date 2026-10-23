//ff:func feature=geulmd type=encoder control=sequence
//ff:what renderImportDoc — IMPORT-DOC as `#id issuer doc [v=version]`
package geulmd

import (
	"strconv"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

// renderImportDoc renders Meta IMPORT-DOC-*: the IMPORT ID, the Issuer (its
// width names the codec), the 4-word document ID and, when pinned, the 2-word
// version ID. An unpinned import (version 0) follows the latest version.
func renderImportDoc(v inspect.ImportDoc) string {
	s := "IMPORT-DOC #" + strconv.FormatUint(v.ID, 10) + " " + core.FormatID(v.Issuer) + " " + core.FormatID(v.Doc)
	if v.Version != 0 {
		s += " v=" + core.FormatID([]core.Word{core.Word(v.Version >> 16), core.Word(v.Version)})
	}
	return s
}

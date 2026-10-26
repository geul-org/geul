//ff:func feature=notation type=decoder control=sequence
//ff:what docAddress — a document as <Issuer>/<문서 ID> in canonical hex, @version when it is not 0
package notation

import (
	"strconv"

	"github.com/geul-org/geul/go/pkg/core"
)

// docAddress writes issuer/doc (GEUL-ORG ID 표기) and `@v` for v ≠ 0.
func docAddress(issuer, doc []core.Word, v uint32) string {
	s := core.FormatID(issuer) + "/" + core.FormatID(doc)
	if v != 0 {
		s += "@" + strconv.FormatUint(uint64(v), 10)
	}
	return s
}

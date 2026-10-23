//ff:func feature=schema type=rule control=sequence
//ff:what applyBind — P5: a version list BIND names a published document of the list's Issuer with a fresh or same-document slug
package schema

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// applyBind checks one version list BIND (P5, edge/PUBLICATION 문서 이름) and
// records the name. The Subject is the document itself in the list's Issuer,
// the document was published before, the name is a slug, and a name already
// bound stays with its document.
func applyBind(st *listState, r rec) []Fact {
	b, err := core.DecodeBind(r.R)
	if err != nil {
		return []Fact{factAt("P5", Error, r, "Payload", "edge/BIND 레이아웃", err.Error())}
	}
	doc, ok := bindDoc(b.Subject, st.issuer)
	if !ok {
		return []Fact{factAt("P5", Error, r, "Subject", "EXTERNAL-ISSUER(목록의 Issuer, 문서 ID, 0)", "다른 Reference")}
	}
	name := string(b.Name)
	if st.docs[doc] == nil {
		return []Fact{factAt("P5", Error, r, "Subject", "앞에 공표된 문서", doc)}
	}
	if !isSlug(name) {
		return []Fact{factAt("P5", Error, r, "이름", "소문자 ASCII 슬러그 (a-z0-9-)", fmt.Sprintf("%q", name))}
	}
	if prev, bound := st.names[name]; bound && prev != doc {
		return []Fact{factAt("P5", Error, r, "이름", "다른 문서에 묶이지 않은 이름", fmt.Sprintf("%q는 %s에 묶임", name, prev))}
	}
	st.names[name] = doc
	return nil
}

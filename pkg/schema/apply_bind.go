//ff:func feature=schema type=rule control=sequence
//ff:what applyBind — P5: a version list BIND names the listed, published document with a slug not bound elsewhere
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// applyBind checks one version list BIND (P5, edge/PUBLICATION 문서 이름) and
// records the name. The Subject is the listed document itself, published
// earlier in the list, and the name is a slug.
func applyBind(st *listState, r rec) []Fact {
	b, err := core.DecodeBind(r.R)
	if err != nil {
		return []Fact{factAt("P5", Error, r, "Payload", "edge/BIND 레이아웃", err.Error())}
	}
	doc, ok := bindDoc(b.Subject)
	if !ok || doc != st.doc {
		return []Fact{factAt("P5", Error, r, "Subject", "EXTERNAL-ISSUER(Issuer, "+st.doc+", 0)", "다른 Reference")}
	}
	if st.docs[doc] == nil {
		return []Fact{factAt("P5", Error, r, "Subject", "앞에 공표된 문서", doc)}
	}
	name := string(b.Name)
	if !isSlug(name) {
		return []Fact{factAt("P5", Error, r, "이름", "소문자 ASCII 슬러그 (a-z0-9-)", fmt.Sprintf("%q", name))}
	}
	st.names[name] = doc
	return nil
}

//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what isVersionList — a version list document: no DOC-ID and a BODY that opens with a PUBLICATION
package schema

import "github.com/geul-org/gogeul/pkg/core"

// isVersionList reports whether the document is a version list
// (edge/PUBLICATION, `<Issuer>/<문서 ID>/versions.geul`): it has no DOC-ID of
// its own and its BODY opens with a PUBLICATION record.
func isVersionList(parts docParts) bool {
	for _, r := range parts.Head {
		if r.R.Codec == core.MetaDocID {
			return false
		}
	}
	return len(parts.Body) > 0 && parts.Body[0].R.Kind == core.KindEdge && parts.Body[0].R.Codec == core.EdgePublication
}

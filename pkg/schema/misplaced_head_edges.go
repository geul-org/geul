//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what misplacedHeadEdges — E8: a SOURCE on the whole document or a QUALIFIER on a HEAD SOURCE found in the BODY
package schema

import "github.com/geul-org/gogeul/pkg/core"

// misplacedHeadEdges reports every BODY record that is a HEAD Edge
// (core.HeadEdge against the HEAD's SOURCEs, headSources): the source of the
// whole document and the QUALIFIERs on it are written in the HEAD only
// (document HEAD와 BODY).
func misplacedHeadEdges(parts docParts) []Fact {
	sources := headSources(parts.Head)
	var fs []Fact
	for _, r := range parts.Body {
		if _, ok := core.HeadEdge(r.R, sources); ok {
			fs = append(fs, factAt("E8", Error, r, "", "문서 전체 SOURCE와 그 QUALIFIER는 HEAD에", "BODY의 "+core.CodecName(core.KindEdge, r.R.Codec)))
		}
	}
	return fs
}

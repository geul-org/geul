//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what ParseDocumentNames — the name → document map of an Issuer's documents.geul (BIND records)
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// ParseDocumentNames reads a documents.geul (GEUL-ORG 문서 이름): in
// `<Issuer>/documents.geul` every BIND whose Subject is
// DOC-POINTER(Issuer, document, 0) maps its name to that document, and a
// Subject that is a codec TAG WORD maps it to that codec's spec document
// (codecKey). Other records are ignored. A
// later BIND of the same name wins.
func ParseDocumentNames(b []byte) (map[string]DocKey, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return nil, err
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return nil, err
	}
	names := map[string]DocKey{}
	for _, r := range doc.Body {
		bind, err := core.DecodeBind(r)
		if err != nil {
			continue
		}
		if d, ok := docOf(bind.Subject); ok {
			names[string(bind.Name)] = d
		}
		if t, ok := core.CodecSubject(bind.Subject); ok {
			names[string(bind.Name)] = codecKey(t)
		}
	}
	return names, nil
}

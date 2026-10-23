//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what importKeys — the Keys a document's IMPORT-DOC-* records ask for, by IMPORT ID
package resolve

import "github.com/geul-org/geul/go/pkg/inspect"

// importKeys maps each IMPORT ID of doc to the Key its IMPORT-DOC-* asks for.
func importKeys(doc inspect.Document) map[uint64]Key {
	out := map[uint64]Key{}
	for _, e := range doc.Head {
		if imp, ok := e.Data.(inspect.ImportDoc); ok {
			out[imp.ID] = Key{Issuer: imp.Issuer, Doc: imp.Doc, Version: imp.Version}
		}
	}
	return out
}

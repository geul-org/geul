//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what docIDOf — the DOC-ID of an interpreted document
package resolve

import "github.com/geul-org/gogeul/pkg/inspect"

// docIDOf returns doc's Meta DOC-ID; false for an unnamed document.
func docIDOf(doc inspect.Document) (inspect.DocID, bool) {
	for _, e := range doc.Head {
		if n, ok := e.Data.(inspect.DocID); ok {
			return n, true
		}
	}
	return inspect.DocID{}, false
}

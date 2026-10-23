//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what docNameOf — the DOC-NAME of an interpreted document
package resolve

import "github.com/geul-org/geul/go/pkg/inspect"

// docNameOf returns doc's Meta DOC-NAME; false for an unnamed document.
func docNameOf(doc inspect.Document) (inspect.DocName, bool) {
	for _, e := range doc.Head {
		if n, ok := e.Data.(inspect.DocName); ok {
			return n, true
		}
	}
	return inspect.DocName{}, false
}

//ff:func feature=page type=builder control=iteration dimension=1
//ff:what sourceText — the text of the BODY's SOURCE record
package page

import "github.com/geul-org/gogeul/pkg/inspect"

// sourceText returns the authored text of the first SOURCE record; false when
// the BODY has none.
func sourceText(body []inspect.Entry) (string, bool) {
	for _, e := range body {
		if s, ok := e.Data.(inspect.Source); ok {
			return string(s.Bytes), true
		}
	}
	return "", false
}

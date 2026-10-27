//ff:func feature=server type=helper control=iteration dimension=1
//ff:what TestAcceptsGeul — the GEUL media type in Accept, alone or in a list; a browser's Accept or none is not
package server

import (
	"net/http/httptest"
	"testing"
)

func TestAcceptsGeul(t *testing.T) {
	cases := map[string]bool{
		MediaType:                                   true,
		"text/html;q=0.5, application/vnd.geul":     true,
		"text/html,application/xhtml+xml,*/*;q=0.8": false,
		"": false,
	}
	for accept, want := range cases {
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Set("Accept", accept)
		if acceptsGeul(r) != want {
			t.Fatalf("%q: want %v", accept, want)
		}
	}
}

//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestParseEntry — off, registry, file://, https:// (geul.org public, others not); http:// and junk fail
package resolve

import (
	"reflect"
	"testing"
)

func TestParseEntry(t *testing.T) {
	cases := map[string]Entry{
		"off":                      {Name: "off", Off: true},
		"registry":                 {Name: "registry", Source: Registry{}, Public: true},
		"file:///srv":              {Name: "file:///srv", Source: Dir{Root: "/srv"}},
		"https://geul.org/":        {Name: "https://geul.org/", Source: HTTP{Base: "https://geul.org"}, Public: true},
		"https://geul.example.com": {Name: "https://geul.example.com", Source: HTTP{Base: "https://geul.example.com"}},
	}
	for s, want := range cases {
		if e, err := parseEntry(s); err != nil || !reflect.DeepEqual(e, want) {
			t.Fatalf("%q: %+v %v", s, e, err)
		}
	}
	for _, s := range []string{"http://x", "file://", "", "geul.org"} {
		if _, err := parseEntry(s); err == nil {
			t.Fatalf("%q: want error", s)
		}
	}
}

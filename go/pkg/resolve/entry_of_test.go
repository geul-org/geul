//ff:func feature=resolve type=resolver control=sequence
//ff:what TestEntryOf — the entry of a listed version; an unlisted one is false
package resolve

import "testing"

func TestEntryOf(t *testing.T) {
	es := []VersionEntry{{Version: 1}, {Version: 2, Deprecated: true}}
	if e, ok := entryOf(es, 2); !ok || !e.Deprecated {
		t.Fatal("v2")
	}
	if _, ok := entryOf(es, 3); ok {
		t.Fatal("v3: want false")
	}
}

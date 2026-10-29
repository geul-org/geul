//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderManifestName — the name it was opened by
package pack

import "testing"

func TestReaderManifestName(t *testing.T) {
	if n := newReader(memStore{}, Options{}, &manifest{}, "abc").ManifestName(); n != "abc" {
		t.Fatal(n)
	}
}

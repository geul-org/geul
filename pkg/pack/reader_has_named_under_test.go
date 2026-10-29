//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderHasNamedUnder — a prefix match in the sorted NAMED table
package pack

import "testing"

func TestReaderHasNamedUnder(t *testing.T) {
	r := newReader(memStore{}, Options{}, &manifest{named: []namedEntry{{"0x0000/documents.geul", Loc{}}, {"issuers.geul", Loc{}}}}, "")
	if !r.hasNamedUnder("0x0000/") || r.hasNamedUnder("0x0001/") || r.hasNamedUnder("z") {
		t.Fatal("hasNamedUnder")
	}
}

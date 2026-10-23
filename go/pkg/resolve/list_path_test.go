//ff:func feature=resolve type=helper control=sequence
//ff:what TestListPath — Issuer/document/ with a trailing slash
package resolve

import (
	"testing"
)

func TestListPath(t *testing.T) {
	if p := listPath(sampleKey(1, 0).DocKey()); p != "0x0000/0xffff000000000001/" {
		t.Fatalf("got %q", p)
	}
}

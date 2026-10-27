//ff:func feature=codecspec type=decoder control=sequence
//ff:what TestFieldDecls — FIELDs by Edge Id with names; other records skipped
package codecspec

import (
	"testing"
)

func TestFieldDecls(t *testing.T) {
	fs := fieldDecls(sampleBody(), map[uint64]string{2: "doc"})
	if len(fs) != 3 || fs[2].Name != "doc" || fs[3].Width != 2 || fs[3].Name != "" {
		t.Fatalf("got %+v", fs)
	}
}

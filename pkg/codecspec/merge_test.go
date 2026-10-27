//ff:func feature=codecspec type=helper control=sequence
//ff:what TestMerge — the receiver wins; the other's extra codecs are added
package codecspec

import (
	"testing"
)

func TestMerge(t *testing.T) {
	m := Spec{1: {{Name: "mine"}}}.Merge(Spec{1: {{Name: "theirs"}}, 2: {{Name: "extra"}}})
	if len(m) != 2 || m[1][0].Name != "mine" || m[2][0].Name != "extra" {
		t.Fatalf("got %+v", m)
	}
}

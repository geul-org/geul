//ff:func feature=page type=builder control=sequence
//ff:what TestHeadRows — each Meta is a `meta` row anchored h<n> with its fields and words
package page

import "testing"

func TestHeadRows(t *testing.T) {
	rows := headRows(fixtureDoc(t).Head)
	if len(rows) != 2 || rows[1].Anchor != "h1" || rows[1].Name != "PROFILE" || rows[1].Class != "meta" || rows[1].Fields[0].Text != `"tangeul-1.0"` || len(rows[1].HeadParts) != 1 {
		t.Fatalf("got %+v", rows)
	}
}

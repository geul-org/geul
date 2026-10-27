//ff:func feature=cli type=helper control=sequence
//ff:what TestLabelIDs — labelled lines (#n or #name) map to their place after next; unlabelled lines map nothing; a repeated label fails
package cli

import "testing"

func TestLabelIDs(t *testing.T) {
	got, err := labelIDs([]string{"#7 TRIPLE", "LABEL", "#edu ALIAS", "#2 QUALIFIER"}, 10)
	if err != nil || len(got) != 3 || got["7"] != 10 || got["edu"] != 12 || got["2"] != 13 {
		t.Fatal(got, err)
	}
	if _, err := labelIDs([]string{"#a TRIPLE", "#a LABEL"}, 1); err == nil {
		t.Fatal("want error for a repeated label")
	}
}

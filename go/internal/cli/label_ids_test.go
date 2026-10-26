//ff:func feature=cli type=helper control=sequence
//ff:what TestLabelIDs — labelled lines map to their place after next; unlabelled and bad labels map nothing
package cli

import "testing"

func TestLabelIDs(t *testing.T) {
	got := labelIDs([]string{"#7 TRIPLE", "LABEL", "#x ALIAS", "#2 QUALIFIER"}, 10)
	if len(got) != 2 || got[7] != 10 || got[2] != 13 {
		t.Fatal(got)
	}
}

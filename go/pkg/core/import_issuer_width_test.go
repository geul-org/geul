//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestImportIssuerWidth — the four IMPORT-DOC widths and zero for other codecs
package core

import "testing"

func TestImportIssuerWidth(t *testing.T) {
	cases := map[uint8]int{MetaImportDoc1Word: 1, MetaImportDoc2Word: 2, MetaImportDoc4Word: 4, MetaImportDoc16Word: 16, 0x64: 0, MetaDocName: 0}
	for c, want := range cases {
		if got := ImportIssuerWidth(c); got != want {
			t.Fatalf("ImportIssuerWidth(0x%02X)=%d want %d", c, got, want)
		}
	}
}

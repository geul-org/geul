//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIssuerWidth — ISSUER widths for Value/Edge and zero for direct codes, Reference and Meta
package core

import "testing"

func TestIssuerWidth(t *testing.T) {
	cases := []struct {
		k    Kind
		c    uint8
		want int
	}{
		{KindEdge, CodecIssuer1Word, 1}, {KindValue, CodecIssuer2Word, 2},
		{KindEdge, CodecIssuer4Word, 4}, {KindValue, CodecIssuer16Word, 16},
		{KindEdge, 0x60, 0}, {KindMeta, CodecIssuer1Word, 0}, {KindReference, CodecIssuer1Word, 0},
	}
	for _, c := range cases {
		if got := IssuerWidth(c.k, c.c); got != c.want {
			t.Fatalf("IssuerWidth(%d,0x%02X)=%d want %d", c.k, c.c, got, c.want)
		}
	}
}

//ff:func feature=page type=helper control=iteration dimension=1
//ff:what TestRefFields — pointer, issuer, issuer/document pointer and declaration splits, and short payloads
package page

import (
	"reflect"
	"testing"
)

func TestRefFields(t *testing.T) {
	cases := []struct {
		c    uint8
		n    int
		want []refField
	}{
		{0x9, 3, []refField{{"IMPORT ID", 2}, {"대상 ID", 1}}},
		{0xC, 6, []refField{{"Issuer", 1}, {"문서 ID", 4}, {"대상 ID", 1}}},
		{0xF, 21, []refField{{"Issuer", 16}, {"문서 ID", 4}, {"대상 ID", 1}}},
		{0xD, 3, []refField{{"Issuer", 2}, {"문서 ID", 1}, {"대상 ID", 0}}},
		{0x2, 2, []refField{{"대상 ID", 2}}},
		{0x6, 2, []refField{{"Issuer", 2}}},
		{0x7, 7, []refField{{"Issuer", 1}, {"문서 ID", 4}, {"버전 ID", 2}}},
		{0x7, 22, []refField{{"Issuer", 16}, {"문서 ID", 4}, {"버전 ID", 2}}},
		{0x7, 5, []refField{{"Issuer", 0}, {"문서 ID", 4}, {"버전 ID", 1}}},
	}
	for i, c := range cases {
		if got := refFields(c.c, c.n); !reflect.DeepEqual(got, c.want) {
			t.Fatalf("case %d: %+v", i, got)
		}
	}
}

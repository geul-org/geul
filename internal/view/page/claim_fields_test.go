//ff:func feature=page type=builder control=sequence
//ff:what TestClaimFields — REF, integer and span objects
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestClaimFields(t *testing.T) {
	f := claimFields(toulmin.Claim{Subject: 1, Predicate: 2, Kind: toulmin.ObjTID, ObjTID: 3, Mode: 4}, nil)
	if f[0].Text != "predicate 0x0002 mode 0x0004" || f[0].Refs[0].ID != 1 || f[1].Refs[0].ID != 3 {
		t.Fatalf("ref: %+v", f)
	}
	if f := claimFields(toulmin.Claim{Kind: toulmin.ObjNum, ObjNum: 9}, nil); f[1].Text != "9" || f[0].Refs != nil {
		t.Fatalf("num: %+v", f)
	}
	if f := claimFields(toulmin.Claim{Kind: toulmin.ObjSpan, ObjSpan: toulmin.Span{Off: 1, Len: 2}}, nil); f[1].Text != "@1+2" {
		t.Fatalf("span: %+v", f)
	}
}

//ff:func feature=geulmd type=encoder control=sequence
//ff:what TestClaimStr — REF, integer and span objects
package geulmd

import (
	"testing"

	"github.com/geul-org/geul/go/pkg/toulmin"
)

func TestClaimStr(t *testing.T) {
	if claimStr(toulmin.Claim{Subject: 1, Predicate: 2, ObjTID: 3, Mode: 4}) != "#1/0x0002/#3/0x0004" {
		t.Fatal("ref object")
	}
	if claimStr(toulmin.Claim{Kind: toulmin.ObjNum, ObjNum: 9}) != "#0/0x0000/n:9/0x0000" {
		t.Fatal("integer object")
	}
	if claimStr(toulmin.Claim{Kind: toulmin.ObjSpan, ObjSpan: toulmin.Span{Off: 1, Len: 2}}) != "#0/0x0000/@1+2/0x0000" {
		t.Fatal("span object")
	}
}

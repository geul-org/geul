//ff:func feature=geulmd type=encoder control=iteration dimension=1
//ff:what TestRender — a full document renders as a HEAD block and a BODY block with every expected line; bad input fails
package geulmd

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	out, err := Render(sampleDoc(t))
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"```geul\nSPEC-VERSION \"geul-t002\"\nPROFILE \"tangeul-1.0\"\nDEFAULT-LANGUAGE lang\"ko\"\nDOC-NAME 0x0001 0x0000000000000005\nIMPORT-DOC #1 0x0001 0x0000000000000002\nRECORD META 0x2a : 0x0007\n```\n",
		"SOURCE #9 span=1 \"# t\\n\"", "TEXT #10 \"n\"", "SECTION #11 5", "CASE-DECLARE #1 role=0 level=0 @0+4",
		"LABEL #12 #1 -> #10 @0+4", "RECORD EDGE 0x61 : 0x00c0 0x000d 0x0000", "VALUE uint(5)", "RECORD VALUE 0x03 : 0x3c00",
	}
	for _, w := range want {
		if !strings.Contains(out, w) {
			t.Fatalf("missing %q in:\n%s", w, out)
		}
	}
	if _, err := Render([]byte{1}); err == nil {
		t.Fatal("odd bytes: want error")
	}
	if _, err := Render([]byte{0x18, 0x0F}); err == nil {
		t.Fatal("not a document: want error")
	}
}

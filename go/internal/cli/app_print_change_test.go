//ff:func feature=cli type=helper control=sequence
//ff:what TestAppPrintChange — -, + and ~ with its → line
package cli

import (
	"bytes"
	"testing"
)

func TestAppPrintChange(t *testing.T) {
	var b bytes.Buffer
	a := &app{out: &b}
	a.printChange(changeLine{Op: "-", Old: "o"})
	a.printChange(changeLine{Op: "+", New: "n"})
	a.printChange(changeLine{Op: "~", Old: "o", New: "n"})
	if b.String() != "- o\n+ n\n~ o\n→ n\n" {
		t.Fatalf("%q", b.String())
	}
}

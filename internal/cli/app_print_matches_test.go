//ff:func feature=cli type=helper control=sequence
//ff:what TestAppPrintMatches — only matching lines, prefixed
package cli

import (
	"bytes"
	"regexp"
	"testing"
)

func TestAppPrintMatches(t *testing.T) {
	var b bytes.Buffer
	(&app{out: &b}).printMatches("d", regexp.MustCompile("a"), []string{"abc", "xyz"})
	if b.String() != "d: abc\n" {
		t.Fatal(b.String())
	}
}

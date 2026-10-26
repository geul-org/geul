//ff:func feature=cli type=helper control=sequence
//ff:what TestAppReadLines — standard input without blanks and comments; a missing file fails
package cli

import (
	"slices"
	"strings"
	"testing"
)

func TestAppReadLines(t *testing.T) {
	a := &app{in: strings.NewReader(" a \n\n// c\nb\n")}
	if got, err := a.readLines("-"); err != nil || !slices.Equal(got, []string{"a", "b"}) {
		t.Fatalf("%q %v", got, err)
	}
	if _, err := a.readLines("/nonexistent/list"); err == nil {
		t.Fatal("want error")
	}
}

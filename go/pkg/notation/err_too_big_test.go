//ff:func feature=notation type=encoder control=sequence
//ff:what TestErrTooBig — the message names the number
package notation

import (
	"strings"
	"testing"
)

func TestErrTooBig(t *testing.T) {
	if !strings.Contains(errTooBig("99").Error(), "99") {
		t.Fatal("message")
	}
}

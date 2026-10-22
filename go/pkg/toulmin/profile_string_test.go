//ff:func feature=toulmin type=encoder control=sequence
//ff:what TestProfileString — the PROFILE text of this build passes CheckProfile
package toulmin

import "testing"

func TestProfileString(t *testing.T) {
	if got := ProfileString(); got != "tangeul-1.0" {
		t.Fatalf("got %q", got)
	}
}

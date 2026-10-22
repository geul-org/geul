//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what TestCheckProfile — same major with any minor passes; other names, majors and junk fail
package toulmin

import "testing"

func TestCheckProfile(t *testing.T) {
	for _, ok := range []string{"tangeul-1.0", "tangeul-1.7"} {
		if err := CheckProfile(ok); err != nil {
			t.Fatalf("%q: %v", ok, err)
		}
	}
	for _, bad := range []string{"tangeul-0.4", "tangle-1.0", "tangeul-1.0x", "tangeul-01.0", ""} {
		if err := CheckProfile(bad); err == nil {
			t.Fatalf("%q must fail", bad)
		}
	}
}

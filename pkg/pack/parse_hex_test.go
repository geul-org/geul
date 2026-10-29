//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParseHex — canonical lowercase 0x hex of the exact width only
package pack

import "testing"

func TestParseHex(t *testing.T) {
	if b, ok := parseHex("0x00ff", 4); !ok || b[1] != 0xff {
		t.Fatal("valid")
	}
	for _, s := range []string{"0x0ff", "0X00ff", "0x00FF", "1x00ff", "0x00fg"} {
		if _, ok := parseHex(s, 4); ok {
			t.Fatal(s)
		}
	}
}

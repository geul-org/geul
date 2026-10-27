//ff:func feature=config type=helper control=sequence
//ff:what TestPick — flag, then environment, then file, then the default
package config

import "testing"

func TestPick(t *testing.T) {
	def := Value{Value: "d", Source: "default"}
	if pick("f", "e", "x", def).Source != "flag" || pick("", "e", "x", def).Source != "env" || pick("", "", "x", def).Source != "file" || pick("", "", "", def) != def {
		t.Fatal("order")
	}
}

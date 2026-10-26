//ff:func feature=config type=helper control=sequence
//ff:what TestSourceOf — flag, environment, default
package config

import "testing"

func TestSourceOf(t *testing.T) {
	if sourceOf("x", "") != "flag" || sourceOf("", "y") != "env" || sourceOf("", "") != "default" {
		t.Fatal("sourceOf")
	}
}

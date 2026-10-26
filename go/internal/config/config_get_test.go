//ff:func feature=config type=helper control=sequence
//ff:what TestConfigGet — a set key's value; an unknown key is empty
package config

import "testing"

func TestConfigGet(t *testing.T) {
	c := Config{values: map[string]Value{"GEULDATA": {Value: "/d", Source: "env"}}}
	if c.Get("GEULDATA") != "/d" || c.Get("NOPE") != "" {
		t.Fatal("get")
	}
}

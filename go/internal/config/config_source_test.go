//ff:func feature=config type=helper control=sequence
//ff:what TestConfigSource — a set key's source; an unknown key has none
package config

import "testing"

func TestConfigSource(t *testing.T) {
	c := Config{values: map[string]Value{"GEULDATA": {Value: "/d", Source: "env"}}}
	if c.Source("GEULDATA") != "env" || c.Source("NOPE") != "" {
		t.Fatal("source")
	}
}

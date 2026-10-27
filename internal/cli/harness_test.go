//ff:type feature=cli type=model
//ff:what harness — test fixture: a data folder, a config folder and an environment to run geul in
package cli

import "testing"

// harness runs geul commands in a temporary data folder (cwd, so it is
// found) with its own config folder and environment.
type harness struct {
	t    *testing.T
	data string
	conf string
	cwd  string
	env  map[string]string
}

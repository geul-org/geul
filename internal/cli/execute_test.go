//ff:func feature=cli type=command control=sequence
//ff:what TestExecute — 0 for a command that runs, 1 for one that fails
package cli

import (
	"os"
	"testing"
)

func TestExecute(t *testing.T) {
	args := os.Args
	defer func() { os.Args = args }()
	os.Args = []string{"geul", "env", "GEULFORMAT"}
	if Execute() != 0 {
		t.Fatal("env: want 0")
	}
	os.Args = []string{"geul", "nope"}
	if Execute() != 1 {
		t.Fatal("unknown command: want 1")
	}
}

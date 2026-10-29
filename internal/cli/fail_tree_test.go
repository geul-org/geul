//ff:type feature=cli type=model
//ff:what failTree — test fixture: a Store that lists its MapFS but fails every ReadFile
package cli

import "testing/fstest"

// failTree lists files it cannot read.
type failTree struct {
	fstest.MapFS
}

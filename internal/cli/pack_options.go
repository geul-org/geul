//ff:func feature=cli type=helper control=sequence
//ff:what packOptions — how the CLI writes pack generations: default sizes, every generation verified before latest.txt moves
package cli

import "github.com/geul-org/gogeul/pkg/pack"

// packOptions are the BuildOptions of every CLI write.
func packOptions() pack.BuildOptions {
	return pack.BuildOptions{Verify: true}
}

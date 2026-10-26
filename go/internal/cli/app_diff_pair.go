//ff:func feature=cli type=helper control=selection
//ff:what app.diffPair — the base and next documents of geul diff's arguments
package cli

import (
	"os"
	"strings"
)

// diffPair reads the two documents newDiffCmd describes. `@v` as the second
// argument is that version of the first argument's document.
func (a *app) diffPair(args []string) ([]byte, []byte, error) {
	first := args[0]
	_, statErr := os.Stat(first)
	switch {
	case len(args) == 2 && strings.HasPrefix(args[1], "@"):
		addr, _, _ := strings.Cut(first, "@")
		return a.readPair(first, addr+args[1])
	case len(args) == 2:
		return a.readPair(first, args[1])
	case statErr == nil:
		return a.draftBase(first)
	case strings.Contains(first, "@"):
		addr, _, _ := strings.Cut(first, "@")
		return a.readPair(first, addr)
	default:
		return a.lastTwo(first)
	}
}

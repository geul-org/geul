//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what synth — n named records of similar, compressible text (a dictionary has something to learn), and their Get
package pack

import (
	"fmt"
	"strings"
)

// synth returns the paths n/0000.geul … and a Get over their bytes.
func synth(n int) ([]string, func(string) ([]byte, error)) {
	data := map[string][]byte{}
	paths := make([]string, n)
	for i := range paths {
		var b strings.Builder
		for j := 0; j < 12; j++ {
			fmt.Fprintf(&b, "<record kind=%q label=\"entity %d part %d\" lang=\"en\" source=\"wikidata dump 2026-10-19\"/>\n", []string{"ALIAS", "LABEL", "CLAIM"}[(i+j)%3], i*7919%1000, j)
		}
		paths[i] = synthPath(i)
		data[paths[i]] = []byte(b.String())
	}
	return paths, memGet(data)
}

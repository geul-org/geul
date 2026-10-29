//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what parseZstd — the options after "zstd:": a level 1–22 and/or nodict, each at most once
package pack

import (
	"fmt"
	"strconv"
)

// parseZstd builds the Zstd of setting s from its options.
func parseZstd(s string, opts []string) (Codec, error) {
	z := &Zstd{}
	for _, o := range opts {
		n, err := strconv.Atoi(o)
		switch {
		case o == "nodict" && z.DictSize == 0:
			z.DictSize = -1
		case err == nil && n >= 1 && n <= 22 && z.Level == 0:
			z.Level = n
		default:
			return nil, fmt.Errorf("geul/pack: codec %q: bad zstd option %q (a level 1–22 or nodict)", s, o)
		}
	}
	return z, nil
}

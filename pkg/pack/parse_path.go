//ff:func feature=pack type=decoder control=sequence
//ff:what parsePath — map an archive path to P, H or the named space (§3.4); non-canonical forms are named
package pack

import (
	"strings"
)

// parsePath maps `<I>/<D>/<V>.geul` and `<I>/<D>/versions.geul` to P,
// `h/0x<H>.geul` to H and everything else, including non-canonical spellings,
// to the named space.
func parsePath(p string) pathRef {
	if h, ok := strings.CutPrefix(p, "h/"); ok {
		if b, ok := parseHex(strings.TrimSuffix(h, ".geul"), 64); ok && strings.HasSuffix(h, ".geul") {
			return pathRef{space: spaceH, hash: string(b)}
		}
	}
	parts := strings.Split(p, "/")
	if len(parts) != 3 || !isIssuer(parts[0]) {
		return pathRef{}
	}
	return parsePPath(parts)
}

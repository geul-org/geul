//ff:func feature=pack type=decoder control=sequence
//ff:what parsePPath — `<I>/<D>/versions.geul` as slot 0 and `<I>/<D>/<V>.geul` (V ≠ 0) as slot V
package pack

import (
	"encoding/binary"
	"strings"
)

// parsePPath maps the three segments of a candidate P path; anything that is
// not a 16-digit document ID with versions.geul or a non-zero 8-digit version
// is named.
func parsePPath(parts []string) pathRef {
	d, ok := parseHex(parts[1], 16)
	if !ok {
		return pathRef{}
	}
	ref := pathRef{space: spaceP, issuer: parts[0], doc: binary.BigEndian.Uint64(d)}
	if parts[2] == "versions.geul" {
		return ref
	}
	v, ok := parseHex(strings.TrimSuffix(parts[2], ".geul"), 8)
	if !ok || !strings.HasSuffix(parts[2], ".geul") || binary.BigEndian.Uint32(v) == 0 {
		return pathRef{}
	}
	ref.slot = binary.BigEndian.Uint32(v)
	return ref
}

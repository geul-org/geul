//ff:func feature=pack type=decoder control=selection
//ff:what ParseCodec — a codec setting (GEULCOMPRESS, --compress): "" follows the set, none is Stored, zstd[:LEVEL][:nodict] is Zstd
package pack

import (
	"fmt"
	"strings"
)

// ParseCodec reads s. "" gives nil (follow the set's codec), "none" Stored,
// "zstd" Zstd at DefaultZstdLevel with a trained dictionary; ":LEVEL" (1–22)
// and ":nodict" adjust it, e.g. "zstd:19:nodict".
func ParseCodec(s string) (Codec, error) {
	parts := strings.Split(strings.ToLower(strings.TrimSpace(s)), ":")
	switch parts[0] {
	case "":
		if len(parts) == 1 {
			return nil, nil
		}
	case "none":
		if len(parts) == 1 {
			return Stored{}, nil
		}
	case "zstd":
		return parseZstd(s, parts[1:])
	}
	return nil, fmt.Errorf("geul/pack: codec %q: want none, zstd, zstd:LEVEL, zstd:nodict or zstd:LEVEL:nodict", s)
}

package schema

import "github.com/geul-org/gogeul/pkg/core"

// fixedValueWidth is the payload width of fixed-width direct Value codecs
// checked by V2 (OVERVIEW Value, docs/value).
var fixedValueWidth = map[uint8]int{
	core.CodecGregorianDate: 2,
	core.CodecPosixSeconds:  4,
	0x23:                    4,  // DURATION-SECONDS
	0x48:                    1,  // LANGUAGE-SHORT-CODE
	0x60:                    8,  // UUID
	0x61:                    16, // SHA-256-DIGEST
	0x62:                    10, // SHA1-DIGEST
}

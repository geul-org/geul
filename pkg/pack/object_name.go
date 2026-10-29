//ff:func feature=pack type=helper control=sequence
//ff:what objectName — `o/<sha256>.gpk` for a data pack, `o/<sha256>.gpx` for an index object
package pack

import "encoding/hex"

// objectName is the store path of o.
func objectName(o object) string {
	ext := ".gpk"
	if o.kind == kindIndex {
		ext = ".gpx"
	}
	return "o/" + hex.EncodeToString(o.hash[:]) + ext
}

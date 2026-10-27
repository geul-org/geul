//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what recordRun — whether words are wholly a run of plausible embedded Reference/Value records
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// recordRun reports whether ws is non-empty and splits exactly into embedded
// records that look intended: a Reference with a target ID, or a Value of a
// codec this library names or an ISSUER codec. Bare data words rarely pass, so
// an unknown layout is not mistaken for records.
func recordRun(ws []core.Word) bool {
	if len(ws) == 0 {
		return false
	}
	for pos := 0; pos < len(ws); {
		r, n, err := core.DecodeRecord(ws[pos:])
		if err != nil || !plausibleEmbedded(r) {
			return false
		}
		pos += n
	}
	return true
}

//ff:func feature=pack type=encoder control=sequence
//ff:what pPath — the archive path of a P key: versions.geul for slot 0, the version file otherwise
package pack

import "fmt"

// pPath is the path of (issuer, doc, slot) relative to `.well-known/geul/`.
func pPath(issuer string, d uint64, s uint32) string {
	if s == 0 {
		return fmt.Sprintf("%s/0x%016x/versions.geul", issuer, d)
	}
	return fmt.Sprintf("%s/0x%016x/0x%08x.geul", issuer, d, s)
}

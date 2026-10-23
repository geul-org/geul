//ff:func feature=schema type=rule control=sequence
//ff:what checkCivilTime — V2: CIVIL-TIME-NANOSECONDS is 4 words within one day
package schema

import "fmt"

// checkCivilTime applies V2 to a CIVIL-TIME-NANOSECONDS Value: 4 words, a signed
// count in [0, 86400×10^9).
func checkCivilTime(r rec) []Fact {
	if fs := checkWidth(r, 4); fs != nil {
		return fs
	}
	p := r.R.Payload
	v := int64(uint64(p[0])<<48 | uint64(p[1])<<32 | uint64(p[2])<<16 | uint64(p[3]))
	if v < 0 || v >= 86400*1e9 {
		return []Fact{factAt("V2", Error, r, "Payload", "[0, 86400×10^9)", fmt.Sprintf("%d", v))}
	}
	return nil
}

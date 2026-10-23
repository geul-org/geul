//ff:func feature=schema type=rule control=sequence
//ff:what checkBool — V2: BOOL is one word, 0x0000 or 0x0001
package schema

import "fmt"

// checkBool applies V2 to a BOOL Value.
func checkBool(r rec) []Fact {
	if fs := checkWidth(r, 1); fs != nil {
		return fs
	}
	if r.R.Payload[0] > 1 {
		return []Fact{factAt("V2", Error, r, "Payload", "0x0000 또는 0x0001", fmt.Sprintf("0x%04X", uint16(r.R.Payload[0])))}
	}
	return nil
}

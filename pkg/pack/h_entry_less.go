//ff:func feature=pack type=writer control=sequence
//ff:what hEntryLess — H run order: key bytes, then object, then offset
package pack

import "bytes"

// hEntryLess orders H entries by key, then by location: of two records with
// the same content hash the later one (higher object, then offset) comes
// last, and the merge keeps the last, as applying the inventories would
// (§7.1: a later record replaces the h/ path).
func hEntryLess(a, b hEntry) bool {
	if c := bytes.Compare(a.key[:], b.key[:]); c != 0 {
		return c < 0
	}
	if a.loc.Obj != b.loc.Obj {
		return a.loc.Obj < b.loc.Obj
	}
	return a.loc.Off < b.loc.Off
}

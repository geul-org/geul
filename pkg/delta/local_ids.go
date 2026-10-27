//ff:type feature=delta type=model
//ff:what LocalIDs — LOCAL ID assignment that carries the previous version's numbers: same statement same ID, new statements from the auto increment
package delta

// LocalIDs hands out LOCAL IDs for a new version of a document so that a
// statement equal to one of the previous version keeps its ID and every new
// one gets the next number (OVERVIEW 개정, meta/lineage/NEXT-LOCAL-ID). A
// deleted ID is never handed out again. Build it with NewLocalIDs.
type LocalIDs struct {
	next  uint64
	edges map[string][]uint64
	defs  map[string][2]uint64
	prev  []uint64
	used  map[uint64]bool
}

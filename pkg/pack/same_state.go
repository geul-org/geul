//ff:func feature=pack type=rule control=iteration dimension=1
//ff:what sameState — ErrCorrupt naming the first path where two path → location maps differ
package pack

// sameState compares the inventory state with the index answer (V3).
func sameState(inv, idx map[string]Loc) error {
	if len(inv) != len(idx) {
		return corrupt("index and inventories disagree on the number of paths")
	}
	for p, l := range inv {
		if got, ok := idx[p]; !ok || got != l {
			return corrupt("index and inventories disagree at " + p)
		}
	}
	return nil
}

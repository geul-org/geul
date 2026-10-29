//ff:func feature=pack type=rule control=iteration dimension=1
//ff:what checkAscending — ErrCorrupt unless keys are strictly ascending
package pack

// checkAscending rejects equal or decreasing neighbours.
func checkAscending(keys []string) error {
	for i := 1; i < len(keys); i++ {
		if keys[i-1] >= keys[i] {
			return corrupt("directory keys not strictly ascending")
		}
	}
	return nil
}

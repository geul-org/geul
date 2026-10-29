//ff:func feature=pack type=rule control=sequence
//ff:what checkChild — a child block must be in its parent's key space and start with the parent's key for it
package pack

// checkChild verifies child i of parent: same space and first key = parent.keys[i].
func checkChild(parent block, i int, child block) error {
	if isPKind(parent.kind) != isPKind(child.kind) || len(child.keys) == 0 || child.keys[0] != parent.keys[i] {
		return corrupt("child block does not match its parent")
	}
	return nil
}

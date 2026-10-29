//ff:func feature=pack type=decoder control=sequence
//ff:what cursor.end — ErrCorrupt when a read failed or bytes are left over
package pack

// end checks that every read succeeded and nothing is left; what names the structure.
func (c *cursor) end(what string) error {
	if c.bad || len(c.b) != 0 {
		return corrupt(what)
	}
	return nil
}

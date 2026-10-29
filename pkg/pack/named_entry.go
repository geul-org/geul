//ff:type feature=pack type=model
//ff:what namedEntry — a NAMED entry: a path outside P and H and its record
package pack

// namedEntry maps a named path to its record.
type namedEntry struct {
	path string
	loc  Loc
}

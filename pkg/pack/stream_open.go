//ff:type feature=pack type=model
//ff:what streamOpen — a Stream's open object: number (-1: none), bytes written, items added
package pack

// streamOpen is an open object: its number (-1 when none is open), the
// bytes written so far and the items added.
type streamOpen struct {
	Num   int
	Off   uint64
	Count uint64
}

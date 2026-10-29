//ff:type feature=pack type=decoder
//ff:what cursor — a little-endian byte reader that turns any overrun into a sticky failure
package pack

// cursor reads b from the front. Once a read runs past the end (or a varint
// is malformed) bad stays true and every later read yields zero values.
type cursor struct {
	b   []byte
	bad bool
}

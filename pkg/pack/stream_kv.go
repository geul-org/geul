//ff:type feature=pack type=model
//ff:what streamKV — a key (P key, H key or named path) and its location, for a checkpoint
package pack

// streamKV is a key (16-byte P key, 32-byte H key or a named path) and its location.
type streamKV struct {
	Key string
	Loc Loc
}

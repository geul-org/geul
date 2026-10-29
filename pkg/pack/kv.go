//ff:type feature=pack type=model
//ff:what kv — one directory entry to write: key and location
package pack

// kv is a key and its location (a deletion marker has Enc 0xFF).
type kv struct {
	key string
	loc Loc
}

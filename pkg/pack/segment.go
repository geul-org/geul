//ff:type feature=pack type=model
//ff:what segment — a stretch of sample bytes chosen for a dictionary, with its score
package pack

// segment is one coverHistory pick; higher scores go nearer the end.
type segment struct {
	score uint64
	b     []byte
}

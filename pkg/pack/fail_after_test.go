//ff:type feature=pack type=model
//ff:what failAfter — a Codec that frames raw for a number of calls, then fails
package pack

// failAfter succeeds while *left > 0 (counting down), then fails.
type failAfter struct{ left *int }

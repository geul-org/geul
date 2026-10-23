//ff:type feature=toulmin type=model
//ff:what testWire — the 1-word SPAN width the packet tests encode with
package toulmin

// testWire is the narrowest document width (1-word SPAN fields), the common
// case for small TANGEUL documents.
var testWire = Wire{Span: 1}

//ff:type feature=notation type=model
//ff:what scanState — where a scan of notation text is: inside a quote, after a backslash, inside <…>, how deep in (…)
package notation

// scanState tracks nesting while scanning notation text so separators inside
// quotes, URIs and call arguments are not split on.
type scanState struct {
	quote  bool
	escape bool
	angle  bool
	depth  int
}

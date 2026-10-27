//ff:func feature=notation type=encoder control=sequence
//ff:what scanState.top — whether the scan is outside every quote, URI and parenthesis
package notation

// top reports whether a separator here splits.
func (s *scanState) top() bool {
	return !s.quote && !s.angle && s.depth == 0
}

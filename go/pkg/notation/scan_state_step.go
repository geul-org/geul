//ff:func feature=notation type=encoder control=selection
//ff:what scanState.step — advance the nesting state over one rune
package notation

// step updates s for rune r: a backslash escapes the next rune inside a
// quote, `"` opens and closes a quote, `<` and `>` an URI outside quotes, `(`
// and `)` a call.
func (s *scanState) step(r rune) {
	switch {
	case s.escape:
		s.escape = false
	case s.quote && r == '\\':
		s.escape = true
	case r == '"':
		s.quote = !s.quote
	case s.quote:
	case r == '<' && !s.angle:
		s.angle = true
	case r == '>' && s.angle:
		s.angle = false
	case s.angle:
	case r == '(':
		s.depth++
	case r == ')':
		s.depth--
	}
}

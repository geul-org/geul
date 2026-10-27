//ff:func feature=notation type=encoder control=sequence
//ff:what flushToken — append the token built so far, if any, and reset the builder
package notation

import "strings"

// flushToken moves cur to out when it is not empty.
func flushToken(out []string, cur *strings.Builder) []string {
	if cur.Len() > 0 {
		out = append(out, cur.String())
		cur.Reset()
	}
	return out
}

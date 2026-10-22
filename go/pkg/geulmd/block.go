//ff:func feature=geulmd type=encoder control=sequence
//ff:what block — wrap lines in a `geul` fenced code block
package geulmd

import "strings"

// block joins lines into a fenced code block with info string `geul`, using a
// fence no backtick run inside can close early.
func block(lines []string) string {
	body := strings.Join(lines, "\n")
	f := fence(body)
	if body != "" {
		body += "\n"
	}
	return f + "geul\n" + body + f + "\n"
}

//ff:func feature=geulmd type=encoder control=iteration dimension=1
//ff:what fence — a backtick fence longer than any backtick run in the block body
package geulmd

import "strings"

// fence returns a backtick fence of at least three backticks and longer than
// every run of backticks in body, so the block can never close early
// (CommonMark fenced code blocks).
func fence(body string) string {
	longest, run := 0, 0
	for _, c := range body {
		if c == '`' {
			run++
		} else {
			run = 0
		}
		if run > longest {
			longest = run
		}
	}
	return strings.Repeat("`", max(3, longest+1))
}

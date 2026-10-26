//ff:func feature=cli type=helper control=selection
//ff:what app.printChange — one change as text: - old, + new, or ~ old then → new
package cli

import "fmt"

// printChange writes c: `- line`, `+ line`, or `~ old` and `→ new`.
func (a *app) printChange(c changeLine) {
	switch c.Op {
	case "-":
		fmt.Fprintln(a.out, "- "+c.Old)
	case "+":
		fmt.Fprintln(a.out, "+ "+c.New)
	default:
		fmt.Fprintln(a.out, "~ "+c.Old)
		fmt.Fprintln(a.out, "→ "+c.New)
	}
}

//ff:func feature=geulmd type=encoder control=iteration dimension=1
//ff:what renderLines — one line per entry in order
package geulmd

import "github.com/geul-org/geul/go/pkg/inspect"

// renderLines renders each entry as one line with line, keeping record order.
func renderLines(entries []inspect.Entry, line func(inspect.Entry) string) []string {
	lines := make([]string, 0, len(entries))
	for _, e := range entries {
		lines = append(lines, line(e))
	}
	return lines
}

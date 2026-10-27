//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what labelIDs — the Edge Id each #label of a batch stands for: its line's place after next; a label used twice fails
package cli

import (
	"fmt"
	"strings"
)

// labelIDs maps the leading `#label` of line i (the text after #, `#7` or
// `#edu`) to next+i. Lines without a label map nothing. A label on two lines
// is an error. The label's form is checked when the line is read
// (notation.ParseEdge).
func labelIDs(lines []string, next uint64) (map[string]uint64, error) {
	out := map[string]uint64{}
	for i, l := range lines {
		first, _, _ := strings.Cut(l, " ")
		label, ok := strings.CutPrefix(first, "#")
		if _, dup := out[label]; ok && dup {
			return nil, fmt.Errorf("line %d: label #%s is used twice", i+1, label)
		}
		if ok {
			out[label] = next + uint64(i)
		}
	}
	return out, nil
}

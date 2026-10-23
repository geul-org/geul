//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what checkWidth — V2: the payload is one of the allowed widths
package schema

import "fmt"

// checkWidth applies V2: the payload width is one of widths.
func checkWidth(r rec, widths ...int) []Fact {
	for _, w := range widths {
		if len(r.R.Payload) == w {
			return nil
		}
	}
	return []Fact{factAt("V2", Error, r, "Payload", fmt.Sprintf("%v WORD", widths), fmtWords(len(r.R.Payload)))}
}

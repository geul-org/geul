//ff:func feature=inspect type=decoder control=sequence
//ff:what parseWidths — Meta ID-WIDTH as Widths; false for a payload that is not one word with the reserved bits clear
package inspect

import "github.com/geul-org/geul/go/pkg/core"

// parseWidths decodes Meta ID-WIDTH `EE VV II` + 10 reserved bits into words.
// ok is false for a payload that is not one word with the reserved bits clear.
func parseWidths(r core.Record) (Widths, bool) {
	if len(r.Payload) != 1 || r.Payload[0]&0x03FF != 0 {
		return Widths{}, false
	}
	w := r.Payload[0]
	return Widths{Edge: core.WidthWords(uint8(w >> 14)), Var: core.WidthWords(uint8(w >> 12)), Import: core.WidthWords(uint8(w >> 10))}, true
}

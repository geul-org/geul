//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what payloadOnly — the payload parts of a layout, without the TAG WORD and its extensions
package notation

import "github.com/geul-org/gogeul/pkg/inspect"

// payloadOnly drops the head parts (TAG WORD, Length, Issuer, format number).
func payloadOnly(parts []inspect.Part) []inspect.Part {
	var out []inspect.Part
	for _, p := range parts {
		if !p.Head {
			out = append(out, p)
		}
	}
	return out
}

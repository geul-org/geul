//ff:func feature=toulmin type=decoder control=sequence
//ff:what DecodeSpan — decode an offset+length Span at the declared SPAN width
package toulmin

import "github.com/geul-org/geul/go/pkg/core"

// DecodeSpan decodes a Span from the head of ws: the offset and then the length,
// each a width-word field.
func DecodeSpan(ws []core.Word, width int) (Span, error) {
	off, err := core.DecodeID(ws, width)
	if err != nil {
		return Span{}, err
	}
	n, err := core.DecodeID(ws[width:], width)
	if err != nil {
		return Span{}, err
	}
	return Span{Off: off, Len: n}, nil
}

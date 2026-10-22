//ff:func feature=toulmin type=encoder control=sequence
//ff:what EncodeSpan — encode a Span as offset then length, each at the declared SPAN width
package toulmin

import "github.com/geul-org/geul/go/pkg/core"

// EncodeSpan encodes s as its offset then its length, each a width-word field
// (GEUL t002 TOULMIN SPAN). A field that does not fit the width is an error.
func EncodeSpan(s Span, width int) ([]core.Word, error) {
	off, err := core.EncodeID(s.Off, width)
	if err != nil {
		return nil, err
	}
	n, err := core.EncodeID(s.Len, width)
	if err != nil {
		return nil, err
	}
	return append(off, n...), nil
}

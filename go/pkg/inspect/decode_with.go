//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what DecodeWith — Decode, laying records out by the document's own codec specs and the given ones
package inspect

import (
	"github.com/geul-org/geul/go/pkg/codecspec"
	"github.com/geul-org/geul/go/pkg/core"
)

// DecodeWith is Decode with codec specs (edge/DECLARE): the specs the document
// declares in its own BODY, then extra for the codecs it does not declare.
// Records of a declared codec carry its fields and are laid out by them.
func DecodeWith(b []byte, extra codecspec.Spec) (Document, error) {
	d, err := Decode(b)
	if err != nil {
		return Document{}, err
	}
	body := make([]core.Record, len(d.Body))
	for i, e := range d.Body {
		body[i] = e.Record
	}
	s := codecspec.Collect(body).Merge(extra)
	applySpec(d.Head, s)
	applySpec(d.Body, s)
	return d, nil
}

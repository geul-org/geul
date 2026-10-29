//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.codecSpecOf — the codec specs one codec spec document's latest version declares
package resolve

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// codecSpecOf collects the BODY of target's latest codec spec document.
func (d Dir) codecSpecOf(target core.Word) (codecspec.Spec, error) {
	b, err := d.latestCodec(target)
	if err != nil {
		return nil, err
	}
	doc, _ := inspect.Decode(b)
	body := make([]core.Record, len(doc.Body))
	for i, e := range doc.Body {
		body[i] = e.Record
	}
	return codecspec.Collect(body), nil
}

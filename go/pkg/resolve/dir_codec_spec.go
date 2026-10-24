//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.CodecSpec — the codec specs of every direct codec spec document published under the GEUL Issuer
package resolve

import (
	"path/filepath"
	"sort"

	"github.com/geul-org/geul/go/pkg/codecspec"
	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

// CodecSpec gathers the specs of the latest version of every codec spec
// document, the GEUL Issuer documents `0x0000/<codecDocID>/` (GEUL-ORG 직접
// 코드 코덱북). Other documents of the Issuer are skipped. The first error
// stops it.
func (d Dir) CodecSpec() (codecspec.Spec, error) {
	lists, _ := filepath.Glob(d.file(core.FormatID([]core.Word{0}) + "/*/" + listFile))
	sort.Strings(lists)
	s := codecspec.Spec{}
	for _, l := range lists {
		target, ok := codecTargetOf(filepath.Base(filepath.Dir(l)))
		if !ok {
			continue
		}
		b, err := d.latestCodec(target)
		if err != nil {
			return nil, err
		}
		doc, _ := inspect.Decode(b)
		body := make([]core.Record, len(doc.Body))
		for i, e := range doc.Body {
			body[i] = e.Record
		}
		s = s.Merge(codecspec.Collect(body))
	}
	return s, nil
}

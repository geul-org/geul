//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.CodecSpec — the codec specs of every direct codec spec document published under proposal/
package resolve

import (
	"path/filepath"
	"sort"

	"github.com/geul-org/geul/go/pkg/codecspec"
	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/inspect"
)

// CodecSpec gathers the specs of the latest version of every codec spec
// document under `proposal/<kind>/<Codec 비트>/` (GEUL-ORG 직접 코드 코덱북).
// The first error stops it.
func (d Dir) CodecSpec() (codecspec.Spec, error) {
	lists, _ := filepath.Glob(d.file("proposal/*/*/" + listFile))
	sort.Strings(lists)
	s := codecspec.Spec{}
	for _, l := range lists {
		target, ok := parseCodecKey(filepath.Base(filepath.Dir(filepath.Dir(l))), filepath.Base(filepath.Dir(l)))
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

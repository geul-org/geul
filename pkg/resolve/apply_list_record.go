//ff:func feature=resolve type=rule control=selection
//ff:what applyListRecord — one version list BODY record: PUBLICATION, DEPRECATION, TOMBSTONE or BIND
package resolve

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// applyListRecord applies one BODY record of the version list of document doc
// (canonical hex). BIND (document names) is kept but changes no version. A
// record of another document, or any other record, is an error.
func applyListRecord(lg *versionLog, r core.Record, doc string) error {
	if r.Kind == core.KindEdge && r.Codec == core.EdgeBind {
		return nil
	}
	p, err := core.DecodePublication(r)
	if err != nil {
		return err
	}
	if core.FormatID(p.Doc) != doc {
		return fmt.Errorf("record for %s in the list of %s", core.FormatID(p.Doc), doc)
	}
	switch p.Codec {
	case core.EdgePublication:
		return addPublication(lg, p)
	default:
		return markVersion(lg, p)
	}
}

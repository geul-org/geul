//ff:func feature=geulmd type=encoder control=selection
//ff:what bodyLine — one BODY line: the typed form of an interpreted record, a raw line otherwise
package geulmd

import (
	"fmt"
	"strconv"

	"github.com/geul-org/geul/go/pkg/inspect"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

// bodyLine renders a BODY entry. A standalone Value shows its literal; a
// TOULMIN record shows its structured form under its spec name. inspect only
// interprets a record when re-encoding gives back its words, so every other
// record — and PREDICATE-DECLARE — is a raw RECORD line.
func bodyLine(e inspect.Entry) string {
	switch p := e.Data.(type) {
	case inspect.Value:
		return "VALUE " + renderValue(p)
	case inspect.Source:
		return fmt.Sprintf("SOURCE #%d span=%d %s", p.TID, p.SpanWidth, strconv.Quote(string(p.Bytes)))
	case toulmin.TextPacket:
		return fmt.Sprintf("TEXT #%d %s", p.TID, strconv.Quote(string(p.Bytes)))
	case toulmin.ControlPacket:
		return controlLine(e.Name, p)
	case toulmin.NodePacket:
		return nodeLine(e.Name, p)
	case toulmin.EdgePacket:
		return linkLine(e.Name, p)
	default:
		return rawRecord(e.Record)
	}
}

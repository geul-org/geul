//ff:func feature=geulmd type=encoder control=selection
//ff:what packetForm — decode a TEXT/control/NODE/LINK record, render its line and re-encode it
package geulmd

import (
	"fmt"
	"strconv"

	"github.com/geul-org/geul/go/pkg/core"
	"github.com/geul-org/geul/go/pkg/toulmin"
)

// packetForm renders a non-SOURCE TOULMIN record of the given class and returns
// the re-encoded words for the lossless check. PREDICATE-DECLARE (whose
// condition layout belongs to the tangl expression AST) and anything else is
// unsupported and reported as an error, so the caller writes a raw line.
func packetForm(name string, ref toulmin.CodecRef, r core.Record, w toulmin.Wire) (string, []core.Word, error) {
	_, _, _, isNode := toulmin.NodeFromCodec(ref)
	_, _, isLink := toulmin.PropFromCodec(ref)
	_, isControl := toulmin.ControlFromFormat(ref.Format)
	switch {
	case ref.Format == toulmin.FormatText:
		p, err := toulmin.DecodeText(r, w)
		enc, _ := toulmin.EncodeText(p, w)
		return fmt.Sprintf("TEXT #%d %s", p.TID, strconv.Quote(string(p.Bytes))), enc, err
	case isControl:
		p, err := toulmin.DecodeControl(r, w)
		if err != nil {
			return "", nil, err
		}
		enc, err := toulmin.EncodeControl(p, w)
		return controlLine(name, p), enc, err
	case isNode:
		p, err := toulmin.DecodeNode(r, w)
		enc, _ := toulmin.EncodeNode(p, w)
		return nodeLine(name, p), enc, err
	case isLink:
		p, err := toulmin.DecodeEdge(r, w)
		enc, _ := toulmin.EncodeEdge(p, w)
		return linkLine(name, p), enc, err
	default:
		return "", nil, fmt.Errorf("geulmd: %s has no structured form", name)
	}
}

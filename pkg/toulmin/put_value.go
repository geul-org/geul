//ff:func feature=toulmin type=encoder control=sequence
//ff:what PutValue — append an embedded Value record
package toulmin

import "github.com/geul-org/gogeul/pkg/core"

// PutValue appends the encoded Value record r to out (GEUL t002: an Edge
// Payload may embed Value records, each self-delimited by its own TAG).
func PutValue(out []core.Word, r core.Record) ([]core.Word, error) {
	ws, err := core.EncodeRecord(r)
	if err != nil {
		return nil, err
	}
	return append(out, ws...), nil
}

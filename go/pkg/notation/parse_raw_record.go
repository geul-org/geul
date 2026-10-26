//ff:func feature=notation type=encoder control=sequence
//ff:what parseRawRecord — a whole record written as raw words: exactly one record, nothing left over
package notation

import (
	"fmt"

	"github.com/geul-org/geul/go/pkg/core"
)

// parseRawRecord reads tok (parseHex) as one encoded record.
func parseRawRecord(tok string) (core.Record, error) {
	ws, err := parseHex(tok)
	if err != nil {
		return core.Record{}, err
	}
	r, n, err := core.DecodeRecord(ws)
	if err != nil || n != len(ws) {
		return core.Record{}, fmt.Errorf("geul/notation: %s is not exactly one record", tok)
	}
	return r, nil
}

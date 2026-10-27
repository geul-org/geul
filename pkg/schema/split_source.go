//ff:func feature=schema type=decoder control=sequence
//ff:what splitSource — a SOURCE Edge cut by its fixed layout: Target and Origin records after the Edge Id, and the words left
package schema

import "github.com/geul-org/gogeul/pkg/core"

// splitSource cuts the payload of r, a SOURCE Edge (edge/SOURCE), after its
// leading Edge Id into the Target record, the Origin record and the words
// left over. It is false when the Edge Id or either record does not decode.
func splitSource(r rec) (target, origin, rest []core.Word, ok bool) {
	_, n, err := core.DecodeRecord(r.R.Payload)
	if err != nil {
		return nil, nil, nil, false
	}
	ws := r.R.Payload[n:]
	_, a, err := core.DecodeRecord(ws)
	if err != nil {
		return nil, nil, nil, false
	}
	_, b, err := core.DecodeRecord(ws[a:])
	if err != nil {
		return nil, nil, nil, false
	}
	return ws[:a], ws[a : a+b], ws[a+b:], true
}

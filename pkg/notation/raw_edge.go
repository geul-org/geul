//ff:func feature=notation type=encoder control=sequence
//ff:what rawEdge — a whole Edge written as raw words: its Edge Id becomes the label, the rest its payload
package notation

import (
	"fmt"
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
)

// rawEdge reads tok as one Edge record opening with a LOCAL Edge Id.
func rawEdge(tok string) (Edge, error) {
	r, err := parseRawRecord(tok)
	if err != nil {
		return Edge{}, err
	}
	lead, n, err := core.DecodeRecord(r.Payload)
	id, lerr := core.DecodeLocalRef(lead)
	if r.Kind != core.KindEdge || err != nil || lerr != nil {
		return Edge{}, fmt.Errorf("geul/notation: %s is not an Edge with a LOCAL Edge Id", tok)
	}
	return Edge{Label: strconv.FormatUint(id, 10), Codec: r.Codec, Rest: r.Payload[n:]}, nil
}

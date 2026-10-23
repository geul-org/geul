//ff:type feature=schema type=model
//ff:what rec — one framed record: number, word offset, total words and the decoded Record
package schema

import "github.com/geul-org/geul/go/pkg/core"

// rec is one record as framing found it: its number, its word offset, its
// total length and the Record split into Kind, Codec, Issuer, Format and Payload.
type rec struct {
	Index int
	Word  int
	Total int
	R     core.Record
}

//ff:func feature=schema type=helper control=sequence
//ff:what edgeOf — test fixture: an Edge with a LOCAL Edge Id followed by words
package schema

import (
	"github.com/geul-org/geul/go/pkg/core"
)

func edgeOf(codec uint8, id uint64, rest ...core.Word) core.Record {
	return core.Record{Kind: core.KindEdge, Codec: codec, Payload: append(enc(core.EncodeLocalRef(id)), rest...)}
}

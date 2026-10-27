//ff:func feature=schema type=helper control=sequence
//ff:what bindOf — test fixture: a BIND Edge with an Edge Id, a Subject and a name
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
)

func bindOf(id uint64, subj core.Record, name string) core.Record {
	return edgeOf(core.EdgeBind, id, append(enc(subj), enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(name)))...)...)
}

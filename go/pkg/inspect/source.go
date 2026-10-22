//ff:type feature=inspect type=model
//ff:what Source — the TOULMIN SOURCE record: the authored text and the SPAN width it declares
package inspect

import "github.com/geul-org/geul/go/pkg/toulmin"

// Source is the TOULMIN SOURCE record: its Edge Id and authored text, and the
// SPAN width in words it declares for the records after it.
type Source struct {
	toulmin.TextPacket
	SpanWidth int
}

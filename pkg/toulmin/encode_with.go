//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what encodeWith — append the F_WITH field: an UNSIGNED-INTEGER count then one TEXT REF per term
package toulmin

import "github.com/geul-org/gogeul/pkg/core"

// encodeWith appends the `with` list: the item count as an embedded
// UNSIGNED-INTEGER Value (at most 4 payload words, so encoding cannot fail),
// then one TEXT REF (LOCAL Reference) per term.
func encodeWith(out []core.Word, with []TID) []core.Word {
	count, _ := core.EncodeRecord(core.EncodeUintValue(uint64(len(with))))
	out = append(out, count...)
	for _, t := range with {
		out = PutID(out, t)
	}
	return out
}

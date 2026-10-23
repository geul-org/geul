//ff:func feature=toulmin type=encoder control=sequence
//ff:what EncodeNode — encode a NodePacket as a NODE record under its type's codec
package toulmin

import (
	"github.com/geul-org/geul/go/pkg/core"
)

// EncodeNode encodes p as a NODE record (GEUL t002 TOULMIN NODE) under the codec
// of its type (rules by role): the Edge Id, the header word (Role, Level, Claim
// object kind, Flags), the flag-gated fields, then the SPAN.
func EncodeNode(p NodePacket, w Wire) ([]core.Word, error) {
	ref, err := NodeCodec(p.Type, p.Role)
	if err != nil {
		return nil, err
	}
	var objKind ObjKind
	if p.Claim != nil {
		objKind = p.Claim.Kind
	}
	header := core.Word(uint16(p.Role)<<14 | uint16(p.Level)<<9 | uint16(objKind)<<7 | uint16(p.Flags))
	body, err := encodeNodeSlots(append(PutID(nil, p.TID), header), p, w)
	if err != nil {
		return nil, err
	}
	body, err = PutSpan(body, p.Span, w.Span)
	if err != nil {
		return nil, err
	}
	return EdgeRecord(ref, body)
}

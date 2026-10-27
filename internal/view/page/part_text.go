//ff:func feature=page type=helper control=selection
//ff:what partText — a declared field as text: hex with its ENUM name, a Value, or a record's codec and payload hex
package page

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// partText shows one field laid out by codec spec (edge/DECLARE 해석 규칙): a
// raw field as its ENUM name, as the Value its CODEC reads it as, as text when it is whole Value or Reference
// records (wholeRecords: a rest field carrying records, such as
// TRANSFORMED-BY's version), or `0x` hex, a record field as its decoded
// Value, else each record it holds by recordText (a rest field may hold
// several, such as TRANSFORMED-BY's URI and version).
func partText(p inspect.Part) string {
	v, isValue := p.Data.(inspect.Value)
	switch {
	case !p.Record && p.Enum != "":
		return p.Enum
	case !p.Record && isValue:
		return valueText(v)
	case !p.Record && wholeRecords(p.Words):
		return recordsText(p.Words)
	case !p.Record:
		return core.FormatID(p.Words)
	case isValue:
		return valueText(v)
	}
	return recordsText(p.Words)
}

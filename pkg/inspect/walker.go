//ff:type feature=inspect type=model
//ff:what walker — a cursor that cuts a payload into Parts field by field, flagging a layout that does not fit
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// walker cuts ws into Parts from pos. sw is the SPAN width in words being
// tried. tail is the width of the fixed raw fields after a codec spec's rest
// field. bad turns true once a field does not fit the remaining words; the
// layout is then rejected. langs names the LANGUAGE-SHORT-CODE language of a
// record field's Value (namedLanguage).
type walker struct {
	ws    []core.Word
	pos   int
	sw    int
	tail  int
	out   []Part
	bad   bool
	langs map[string]string
}

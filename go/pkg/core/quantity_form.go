//ff:func feature=core type=decoder control=selection
//ff:what quantityForm — whether the four QUANTITY records have their forms: SIGNED-DECIMAL amount, Reference or Absent unit, SIGNED-DECIMAL or Absent bounds
package core

import "errors"

// quantityForm checks the forms of a QUANTITY's amount, unit, lower and upper
// records (value/numeric/QUANTITY).
func quantityForm(rs [4]Record) error {
	decimal := func(x Record) bool { return x.Kind == KindValue && x.Codec == CodecSignedDecimal }
	switch {
	case !decimal(rs[0]):
		return errors.New("geul/core: QUANTITY amount is not a SIGNED-DECIMAL Value")
	case rs[1].Kind != KindReference:
		return errors.New("geul/core: QUANTITY unit is not a Reference or Absent")
	case !decimal(rs[2]) && !IsAbsent(rs[2]), !decimal(rs[3]) && !IsAbsent(rs[3]):
		return errors.New("geul/core: QUANTITY bound is not a SIGNED-DECIMAL Value or Absent")
	default:
		return nil
	}
}

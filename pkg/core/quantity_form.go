//ff:func feature=core type=decoder control=selection
//ff:what quantityForm — whether the four QUANTITY records have their forms: SIGNED-DECIMAL amount, Reference or Absent unit, SIGNED-DECIMAL or Absent bounds
package core

import "errors"

// quantityForm checks the forms of a QUANTITY's amount, unit, lower and upper
// records (value/numeric/QUANTITY). The unit is a Reference of a real Codec
// or ABSENT; no other slot marker (Reference Codec 0000) is a unit.
func quantityForm(rs [4]Record) error {
	decimal := func(x Record) bool { return x.Kind == KindValue && x.Codec == CodecSignedDecimal }
	switch {
	case !decimal(rs[0]):
		return errors.New("geul/core: QUANTITY amount is not a SIGNED-DECIMAL Value")
	case rs[1].Kind != KindReference, rs[1].Codec>>3 == 0 && !IsAbsent(rs[1]):
		return errors.New("geul/core: QUANTITY unit is not a Reference or Absent")
	case !decimal(rs[2]) && !IsAbsent(rs[2]), !decimal(rs[3]) && !IsAbsent(rs[3]):
		return errors.New("geul/core: QUANTITY bound is not a SIGNED-DECIMAL Value or Absent")
	default:
		return nil
	}
}

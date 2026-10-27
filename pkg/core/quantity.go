//ff:type feature=core type=model
//ff:what Quantity — the four records of a QUANTITY Value: amount, unit, lower and upper bound (Absent when left out)
package core

// Quantity is the body of a QUANTITY Value (value/numeric/QUANTITY): the
// amount as a SIGNED-DECIMAL Value, the unit as a Reference, and the lower
// and upper bounds as SIGNED-DECIMAL Values. A unit or bound that is left out
// (no unit, no bounds) is the Absent record (IsAbsent).
type Quantity struct {
	Amount Record
	Unit   Record
	Lower  Record
	Upper  Record
}

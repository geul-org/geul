//ff:type feature=core type=model
//ff:what FieldRest — the FIELD width of a raw field running to the end of the record
package core

// FieldRest is the FIELD width 0xFFFF: a raw field of every word left in the
// record, bounded by the record's own TAG Length; only the last field of a
// DECLARE (GEUL t002 docs/edge/FIELD).
const FieldRest Word = 0xFFFF

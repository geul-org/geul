//ff:type feature=core type=model
//ff:what Coordinate — a decoded COORDINATE-EARTH or COORDINATE-GEO: globe (none for Earth), latitude, longitude, uncertainty (binary32) and its unit
package core

// Coordinate is a point decoded from COORDINATE-EARTH (Globe nil: WGS 84
// Earth) or COORDINATE-GEO (Globe is the body's Reference): latitude and
// longitude in degrees, and the uncertainty radius (binary32) in the unit
// Unit names — a Reference, a Wikidata unit item by preference (Q28390
// degree, Q209426 arcminute, Q829073 arcsecond, Q11573 metre …). Unit nil is
// the slot marker ABSENT: the uncertainty is not known, and Uncertainty is
// then +0. An Uncertainty of 0 with a unit is an exact point.
type Coordinate struct {
	Globe       *Record
	Lat         float64
	Lon         float64
	Uncertainty float32
	Unit        *Record
}

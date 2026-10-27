//ff:type feature=core type=model
//ff:what Coordinate — a decoded COORDINATE-EARTH or COORDINATE-GEO: globe (none for Earth), latitude, longitude, uncertainty
package core

// Coordinate is a point decoded from COORDINATE-EARTH (Globe nil: WGS 84
// Earth) or COORDINATE-GEO (Globe is the body's Reference):
// latitude and longitude in degrees and the uncertainty radius in millimetres
// (0 exact, UncertaintyUnknown unknown).
type Coordinate struct {
	Globe *Record
	Lat   float64
	Lon   float64
	U     uint32
}

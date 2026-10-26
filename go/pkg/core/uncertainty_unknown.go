//ff:type feature=core type=model
//ff:what UncertaintyUnknown — the COORDINATE-* uncertainty that says the radius is not known
package core

// UncertaintyUnknown is the COORDINATE-EARTH / COORDINATE-GEO uncertainty
// meaning "not known" (RFC 5870 without u); 0 means an exact point.
const UncertaintyUnknown uint32 = 0xFFFFFFFF

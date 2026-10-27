//ff:type feature=schema type=model
//ff:what placeState — where the structure pass is: HEAD, BODY or after END, with the HEAD SOURCE Edge Ids
package schema

// placeState tracks the structure pass: whether BODY has begun, whether END
// has been seen, and the Edge Ids of the SOURCEs the HEAD has so far
// (headEdgeRec), on which a QUALIFIER is a HEAD record.
type placeState struct {
	inBody  bool
	ended   bool
	sources map[uint64]bool
}

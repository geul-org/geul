//ff:type feature=schema type=model
//ff:what placeState — where the structure pass is: HEAD, BODY or after END
package schema

// placeState tracks the structure pass: whether BODY has begun and whether END
// has been seen.
type placeState struct {
	inBody bool
	ended  bool
}

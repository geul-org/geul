//ff:type feature=notation type=model
//ff:what recordFieldAny — a record field that allows any codec
package notation

import "github.com/geul-org/gogeul/pkg/codecspec"

// recordFieldAny is the field of a token read with no declared layout.
var recordFieldAny = codecspec.Field{Name: "record"}

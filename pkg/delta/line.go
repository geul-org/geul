//ff:type feature=delta type=model
//ff:what Line — one Edge of a document's last version with the version that added it and the one that last changed it
package delta

import "github.com/geul-org/gogeul/pkg/core"

// Line is one BODY Edge of the last version given to Blame: Added is the first
// version holding its Edge Id, Changed the last version that changed its
// content (Added when never changed).
type Line struct {
	ID      uint64
	Record  core.Record
	Added   uint32
	Changed uint32
}

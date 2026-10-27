//ff:type feature=cli type=model
//ff:what draft — a document being edited: its HEAD and BODY records
package cli

import "github.com/geul-org/gogeul/pkg/core"

// draft is a working file (README 작성 명령): edited by new, add, def, set,
// rm and meta, then published.
type draft struct {
	head []core.Record
	body []core.Record
}

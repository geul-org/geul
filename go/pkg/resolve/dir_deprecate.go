//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.Deprecate — mark one published version deprecated: still served when asked for, no longer the newest
package resolve

import "github.com/geul-org/geul/go/pkg/core"

// Deprecate appends a DEPRECATION of version v of document doc of issuer to
// its version list (mark). A deprecated version is still served by its
// version ID but is no longer the newest (edge/PUBLICATION).
func (d Dir) Deprecate(issuer, doc []core.Word, v uint32) error {
	_, err := d.mark(issuer, doc, v, core.EdgeDeprecation, 0)
	return err
}

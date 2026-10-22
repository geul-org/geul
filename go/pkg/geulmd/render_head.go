//ff:func feature=geulmd type=encoder control=iteration dimension=1
//ff:what renderHead — the HEAD block lines, after a first pass that reads the declared widths
package geulmd

import "github.com/geul-org/geul/go/pkg/core"

// renderHead renders each HEAD Meta record as one line in order. The declared
// widths are read first (ID-WIDTH may follow an IMPORT-DOC that depends on it),
// and the returned context carries them and the profile into BODY rendering.
func renderHead(head []core.Record) ([]string, headCtx) {
	ctx := headCtx{idWidth: 1, importWidth: 1}
	for _, r := range head {
		if r.Codec == core.MetaIDWidth {
			renderIDWidth(r, &ctx)
		}
	}
	lines := make([]string, 0, len(head))
	for _, r := range head {
		lines = append(lines, renderMeta(r, &ctx))
	}
	return lines, ctx
}

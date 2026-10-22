//ff:type feature=geulmd type=model
//ff:what headCtx — the HEAD settings BODY rendering needs: id widths and whether the TOULMIN profile applies
package geulmd

// headCtx carries what the HEAD declares for rendering the BODY: the Edge Id and
// IMPORT ID widths (Meta ID-WIDTH, 1 word when absent) and whether PROFILE
// names the TANGEUL profile, whose records render structurally.
type headCtx struct {
	idWidth     int
	importWidth int
	tangeul     bool
}

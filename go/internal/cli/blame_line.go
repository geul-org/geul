//ff:type feature=cli type=model
//ff:what blameLine — one line of geul blame
package cli

// blameLine is an Edge's line with the versions that added and last changed
// it.
type blameLine struct {
	Added   uint32 `json:"added"`
	Changed uint32 `json:"changed"`
	Line    string `json:"line"`
}

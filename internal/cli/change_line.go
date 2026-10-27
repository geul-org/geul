//ff:type feature=cli type=model
//ff:what changeLine — one change of geul diff with its records as notation lines
package cli

// changeLine is a delta.Change written out: Op "+", "-" or "~", whether it
// is a HEAD record, the Edge Id, and the old and new lines ("" where none).
type changeLine struct {
	Op   string `json:"op"`
	Head bool   `json:"head,omitempty"`
	ID   uint64 `json:"id,omitempty"`
	Old  string `json:"old,omitempty"`
	New  string `json:"new,omitempty"`
}

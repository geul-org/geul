//ff:type feature=cli type=model
//ff:what logJSON — geul log as JSON: the versions and NEXT-LOCAL-ID
package cli

// logJSON is a document's history (resolve.History) for --format json.
type logJSON struct {
	Versions    []logVersion `json:"versions"`
	NextLocalID uint64       `json:"next_local_id"`
}

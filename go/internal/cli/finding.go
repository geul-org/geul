//ff:type feature=cli type=model
//ff:what finding — one problem geul check found: where and what
package cli

// finding is a schema Fact or a version problem of one path.
type finding struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

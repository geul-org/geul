//ff:type feature=cli type=model
//ff:what docLines — a document as notation lines: HEAD and BODY
package cli

// docLines is what show prints: one notation line per record.
type docLines struct {
	Head []string `json:"head"`
	Body []string `json:"body"`
}

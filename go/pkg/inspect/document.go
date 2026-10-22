//ff:type feature=inspect type=model
//ff:what Document — an interpreted GEUL document: HEAD and BODY entries plus the declared widths and profile
package inspect

// Document is one interpreted GEUL document. Head and Body hold one Entry per
// record in order (the DOC, BODY and END structure records are not kept).
// Widths are the ID widths the HEAD declares (1 word each when absent), Profile
// the Meta PROFILE text ("" when absent) and Toulmin whether that profile is the
// TANGEUL profile this build reads, under which TOULMIN records are decoded.
type Document struct {
	Head    []Entry
	Body    []Entry
	Widths  Widths
	Profile string
	Toulmin bool
}

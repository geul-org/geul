//ff:type feature=inspect type=model
//ff:what Document — an interpreted GEUL document: HEAD and BODY entries plus the profile
package inspect

// Document is one interpreted GEUL document. Head and Body hold one Entry per
// record in order (the DOC, BODY and END structure records are not kept).
// Profile is the Meta PROFILE text ("" when absent) and Toulmin whether that profile is the
// TANGEUL profile this build reads, under which TOULMIN records are decoded.
type Document struct {
	Head    []Entry
	Body    []Entry
	Profile string
	Toulmin bool
}

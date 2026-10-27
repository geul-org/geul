//ff:func feature=schema type=decoder control=sequence
//ff:what bindDoc — the document a BIND Subject names: DOC-POINTER(Issuer, document, version 0)
package schema

import "github.com/geul-org/gogeul/pkg/core"

// bindDoc reads a version list BIND Subject, which must be a DOC-POINTER
// Reference with version ID 0 (the document, not one version), and returns
// its document ID in canonical hex.
func bindDoc(subj core.Record) (string, bool) {
	_, doc, ver, err := core.DecodeDocPointer(subj)
	if err != nil || ver[0]|ver[1] != 0 {
		return "", false
	}
	return core.FormatID(doc), true
}

//ff:type feature=schema type=model
//ff:what docParts — the HEAD and BODY records found by the structure pass
package schema

// docParts is the document split by the structure pass. The DOC, BODY and END
// records are not included.
type docParts struct {
	Head []rec
	Body []rec
}

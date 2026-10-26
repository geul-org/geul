//ff:type feature=delta type=model
//ff:what Version — one version of a document for Blame: its version ID and snapshot bytes
package delta

// Version is version ID V of a document and its snapshot Doc.
type Version struct {
	V   uint32
	Doc []byte
}

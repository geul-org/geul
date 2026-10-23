//ff:type feature=toulmin type=model
//ff:what TID — a document-local Edge Id (GEUL t002 LOCAL ID)
package toulmin

// TID is a document-local Edge Id (GEUL t002 OVERVIEW Edge, TOULMIN 공통 필드).
// Every record of a TANGEUL document carries one, issued from 1 in declaration
// order, and every in-document reference is one. On the wire it is an embedded
// LOCAL Reference record in minimal words, so the in-memory type is 64 bits and
// the packet layer never assumes a fixed width.
type TID uint64
